package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "reservation-service/pkg/gen"
)

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
