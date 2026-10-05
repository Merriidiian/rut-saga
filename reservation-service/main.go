package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"net/http"
	"os"
	pb "reservation-service/pkg/gen"
	"sync"
	"time"
)

type server struct {
	pb.UnimplementedReservationServiceServer
	mu     sync.Mutex
	data   map[string]*pb.Reservation
	expiry map[string]time.Time
}

func (s *server) CreateReservation(ctx context.Context, r *pb.CreateReservationRequest) (*pb.Reservation, error) {
	if err := r.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if old, ok := s.data[r.ReservationId]; ok {
		if old.ProductId != r.ProductId || old.CustomerId != r.CustomerId || old.Quantity != r.Quantity {
			return nil, status.Error(codes.AlreadyExists, "Reservation payload differs")
		}
		return proto.Clone(old).(*pb.Reservation), nil
	}
	v := &pb.Reservation{ReservationId: r.ReservationId, ProductId: r.ProductId, CustomerId: r.CustomerId, Quantity: r.Quantity, Status: "Active"}
	s.data[r.ReservationId] = v
	s.expiry[r.ReservationId] = time.Now().Add(time.Duration(r.ExpirationMinutes) * time.Minute)
	log.Printf("Reservation created: %s", r.ReservationId)
	return proto.Clone(v).(*pb.Reservation), nil
}
func (s *server) change(r *pb.ReservationRequest, next string) (*pb.Reservation, error) {
	if err := r.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.data[r.ReservationId]
	if !ok {
		if next == "Cancelled" {
			v = &pb.Reservation{ReservationId: r.ReservationId, Status: "Cancelled"}
			s.data[r.ReservationId] = v
			return proto.Clone(v).(*pb.Reservation), nil
		}
		return nil, status.Error(codes.NotFound, "Reservation not found")
	}
	if v.Status == "Active" && time.Now().After(s.expiry[r.ReservationId]) {
		v.Status = "Expired"
	}
	if next == "Confirmed" && v.Status != "Active" && v.Status != "Confirmed" {
		return nil, status.Error(codes.FailedPrecondition, "Reservation is not active")
	}
	if next != "" {
		v.Status = next
		log.Printf("Reservation %s: %s", r.ReservationId, next)
	}
	return proto.Clone(v).(*pb.Reservation), nil
}
func (s *server) GetReservation(ctx context.Context, r *pb.ReservationRequest) (*pb.Reservation, error) {
	return s.change(r, "")
}
func (s *server) ConfirmReservation(ctx context.Context, r *pb.ReservationRequest) (*pb.Reservation, error) {
	return s.change(r, "Confirmed")
}
func (s *server) CancelReservation(ctx context.Context, r *pb.ReservationRequest) (*pb.Reservation, error) {
	return s.change(r, "Cancelled")
}
func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	g := grpc.NewServer()
	pb.RegisterReservationServiceServer(g, &server{data: make(map[string]*pb.Reservation), expiry: make(map[string]time.Time)})
	go func() {
		if err := g.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	mux := runtime.NewServeMux()
	err = pb.RegisterReservationServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatal(err)
	}
	h := http.NewServeMux()
	h.Handle("/api/", mux)
	h.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("swagger"))))
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("gRPC :50051 HTTP :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, h))
}
