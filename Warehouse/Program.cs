using EventStore.Client;
using Grpc.Net.Client;
using Reservation.V1;
using Warehouse;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddSingleton(new EventStoreClient(EventStoreClientSettings.Create(
    builder.Configuration.GetConnectionString("EventStore")!)));
builder.Services.AddSingleton(new ReservationService.ReservationServiceClient(
    GrpcChannel.ForAddress(builder.Configuration["ReservationUrl"]!)));
builder.Services.AddSingleton<OrderRepository>();
builder.Services.AddScoped<OrderSaga>();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
var app = builder.Build();
app.UseSwagger();
app.UseSwaggerUI();
app.MapPost("/orders", async (CreateOrderCommand command, OrderSaga saga, CancellationToken ct) =>
{
    try
    {
        var result = await saga.ExecuteAsync(command, ct);
        return result.Success ? Results.Created($"/orders/{result.OrderId}", result) : Results.BadRequest(result);
    }
    catch (ArgumentException e) { return Results.BadRequest(new { message = e.Message }); }
});
app.MapGet("/orders/{id:guid}", async (Guid id, OrderRepository repo, CancellationToken ct) =>
{
    var order = await repo.LoadAsync(id, ct);
    return order == null ? Results.NotFound() : Results.Ok(new
    {
        order.Id, order.CustomerId, order.ProductId, order.Quantity,
        order.Status, order.ReservationId, order.Version
    });
});
app.MapGet("/orders/{id:guid}/events", async (Guid id, EventStoreClient client, CancellationToken ct) =>
{
    var read = client.ReadStreamAsync(Direction.Forwards, $"order-{id}", StreamPosition.Start, cancellationToken: ct);
    if (await read.ReadState == ReadState.StreamNotFound) return Results.NotFound();
    var result = new List<object>();
    await foreach (var e in read) result.Add(new { type = e.Event.EventType,
        revision = e.Event.EventNumber.ToUInt64(), data = OrderRepository.Decode(e.Event) });
    return Results.Ok(result);
});
app.Run();
