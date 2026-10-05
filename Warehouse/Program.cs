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

app.MapApiEndpoints();
app.Run();
