namespace Warehouse;

public record OrderCreated(Guid Id, Guid CustomerId, Guid ProductId, int Quantity);
public record OrderConfirmed(Guid Id, string ReservationId);
public record OrderCancelled(Guid Id);
public class Order
{
    public Guid Id
    {
        get; private set;
    }
    public Guid CustomerId
    {
        get; private set;
    }
    public Guid ProductId
    {
        get; private set;
    }
    public int Quantity
    {
        get; private set;
    }
    public string Status { get; private set; } = "Created";
    public string? ReservationId
    {
        get; private set;
    }
    public long Version { get; private set; } = -1;
    public List<object> Changes { get; } = new();

    public static Order Create(Guid id, Guid customerId, Guid productId, int quantity)
    {
        if (customerId == Guid.Empty || productId == Guid.Empty || quantity <= 0)
        {
            throw new ArgumentException("Invalid order");
        }

        var order = new Order();
        order.Raise(new OrderCreated(id, customerId, productId, quantity));
        return order;
    }

    public void Confirm(string reservationId)
    {
        if (Status != "Created")
        {
            throw new InvalidOperationException("Order is not new");
        }

        Raise(new OrderConfirmed(Id, reservationId));
    }

    public void Cancel()
    {
        if (Status == "Cancelled")
        {
            return;
        }

        Raise(new OrderCancelled(Id));
    }

    private void Raise(object value)
    {
        Apply(value);
        Changes.Add(value);
    }

    public void Apply(object value)
    {
        switch (value)
        {
            case OrderCreated e:
                Id = e.Id;
                CustomerId = e.CustomerId;
                ProductId = e.ProductId;
                Quantity = e.Quantity;
                Status = "Created";
                break;
            case OrderConfirmed e:
                Status = "Confirmed";
                ReservationId = e.ReservationId;
                break;
            case OrderCancelled:
                Status = "Cancelled";
                break;
        }
        Version++;
    }

}
