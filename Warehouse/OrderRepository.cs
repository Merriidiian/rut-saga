using EventStore.Client;
using System.Text.Json;

namespace Warehouse;

public class OrderRepository(EventStoreClient client)
{
    public static object Decode(EventRecord e) => e.EventType switch
    {
        nameof(OrderCreated) => JsonSerializer.Deserialize<OrderCreated>(e.Data.Span)!,
        nameof(OrderConfirmed) => JsonSerializer.Deserialize<OrderConfirmed>(e.Data.Span)!,
        nameof(OrderCancelled) => JsonSerializer.Deserialize<OrderCancelled>(e.Data.Span)!,
        _ => throw new InvalidOperationException(e.EventType)
    };

    public async Task SaveAsync(Order order, CancellationToken ct = default)
    {
        if (order.Changes.Count == 0)
        {
            return;
        }

        var expected = order.Version - order.Changes.Count;
        var data = order.Changes.Select(e => new EventData(Uuid.NewUuid(),
            e.GetType().Name, JsonSerializer.SerializeToUtf8Bytes(e, e.GetType()))).ToArray();
        if (expected < 0)
        {
            await client.AppendToStreamAsync($"order-{order.Id}", StreamState.NoStream,
                data, cancellationToken: ct);
        }
        else
        {
            await client.AppendToStreamAsync($"order-{order.Id}", new StreamRevision((ulong)expected),
                data, cancellationToken: ct);
        }

        order.Changes.Clear();
    }

    public async Task<Order?> LoadAsync(Guid id, CancellationToken ct = default)
    {
        var order = new Order();
        var events = client.ReadStreamAsync(Direction.Forwards, $"order-{id}",
            StreamPosition.Start, cancellationToken: ct);
        if (await events.ReadState == ReadState.StreamNotFound)
        {
            return null;
        }

        await foreach (var e in events)
        {
            order.Apply(Decode(e.Event));
        }

        return order;
    }
}
