package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "go_grpc_service/go_grpc_service"
)

type server struct {
	pb.UnimplementedGRPCServiceServer
}

func (s *server) SendFeedback(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	// Handle the received feedback message and return a response
	message := req.GetMessage()
	response := &pb.Response{
		Message: fmt.Sprintf("Received feedback: %s", message + "xxx ok xxx"),
	}
	return response, nil
}

func main() {
	// Create a TCP listener on a specific port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	srv := grpc.NewServer()

	// Register the service implementation
	pb.RegisterGRPCServiceServer(srv, &server{})

	// Start the server and listen for incoming connections
	log.Println("Starting gRPC server...")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

