package main

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	pb "reservation-service/pkg/gen"
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
	reservation := &pb.Reservation{
		ReservationId: r.ReservationId,
		ProductId:     r.ProductId,
		CustomerId:    r.CustomerId,
		Quantity:      r.Quantity,
		Status:        "Active",
	}
	s.data[r.ReservationId] = reservation
	s.expiry[r.ReservationId] = time.Now().Add(time.Duration(r.ExpirationMinutes) * time.Minute)
	log.Printf("Reservation created: %s", r.ReservationId)
	return proto.Clone(reservation).(*pb.Reservation), nil
}

func (s *server) change(r *pb.ReservationRequest, next string) (*pb.Reservation, error) {
	if err := r.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	reservation, ok := s.data[r.ReservationId]
	if !ok {
		if next == "Cancelled" {
			reservation = &pb.Reservation{ReservationId: r.ReservationId, Status: "Cancelled"}
			s.data[r.ReservationId] = reservation
			return proto.Clone(reservation).(*pb.Reservation), nil
		}
		return nil, status.Error(codes.NotFound, "Reservation not found")
	}
	if reservation.Status == "Active" && time.Now().After(s.expiry[r.ReservationId]) {
		reservation.Status = "Expired"
	}
	if next == "Confirmed" && reservation.Status != "Active" && reservation.Status != "Confirmed" {
		return nil, status.Error(codes.FailedPrecondition, "Reservation is not active")
	}
	if next != "" {
		reservation.Status = next
		log.Printf("Reservation %s: %s", r.ReservationId, next)
	}
	return proto.Clone(reservation).(*pb.Reservation), nil
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
