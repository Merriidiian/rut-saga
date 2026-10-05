using Reservation.V1;

namespace Warehouse;

public record CreateOrderCommand(Guid CustomerId, Guid ProductId, int Quantity,
    bool FailAfterReservation = false);
public record SagaResult(bool Success, Guid OrderId, string? ReservationId,
    string Message, bool CompensationExecuted);

public class OrderSaga(OrderRepository repository,
    ReservationService.ReservationServiceClient client, ILogger<OrderSaga> logger)
{
    public async Task<SagaResult> ExecuteAsync(CreateOrderCommand command, CancellationToken ct)
    {
        var id = Guid.NewGuid();
        var order = Order.Create(id, command.CustomerId, command.ProductId, command.Quantity);
        var reservationId = id.ToString();
        logger.LogInformation("SAGA {Id} STARTED", id);
        try
        {
            await repository.SaveAsync(order, ct);
            logger.LogInformation("SAGA {Id} OrderCreated saved", id);
            await client.CreateReservationAsync(new CreateReservationRequest
            {
                ReservationId = reservationId,
                ProductId = command.ProductId.ToString(),
                CustomerId = command.CustomerId.ToString(),
                Quantity = command.Quantity,
                ExpirationMinutes = 30
            }, deadline: DateTime.UtcNow.AddSeconds(5), cancellationToken: ct);
            logger.LogInformation("SAGA {Id} Reservation created", id);
            if (command.FailAfterReservation)
            {
                throw new InvalidOperationException("Test failure after reservation");
            }

            await client.ConfirmReservationAsync(new ReservationRequest { ReservationId = reservationId },
                deadline: DateTime.UtcNow.AddSeconds(5), cancellationToken: ct);
            order.Confirm(reservationId);
            await repository.SaveAsync(order, ct);
            logger.LogInformation("SAGA {Id} COMPLETED", id);
            return new(true, id, reservationId, "Order confirmed", false);
        }
        catch (Exception error)
        {
            logger.LogWarning("SAGA {Id} COMPENSATION STARTED: {Message}", id, error.Message);
            var compensated = true;
            try
            {
                await client.CancelReservationAsync(new ReservationRequest { ReservationId = reservationId },
                    deadline: DateTime.UtcNow.AddSeconds(5));
            }
            catch (Exception e)
            {
                compensated = false;
                logger.LogError("SAGA {Id} Reservation cancellation failed: {Message}", id, e.Message);
            }
            try
            {
                var saved = await repository.LoadAsync(id, ct: CancellationToken.None);
                if (saved != null)
                {
                    saved.Cancel();
                    await repository.SaveAsync(saved);
                }
            }
            catch (Exception e)
            {
                compensated = false;
                logger.LogError("SAGA {Id} Order cancellation failed: {Message}", id, e.Message);
            }
            logger.LogWarning("SAGA {Id} COMPENSATION {Status}", id, compensated ? "COMPLETED" : "INCOMPLETE");
            return new(false, id, reservationId, error.Message, compensated);
        }
    }
}
