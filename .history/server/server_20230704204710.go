package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "go_grpc_service/go_grpc_service"
)

type server struct{}

func (s *server) SendFeedback(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	// Handle the received feedback message and return a response
	message := req.GetMessage()
	response := &pb.Response{
		Message: fmt.Sprintf("Received feedback: %s", message),
	}
	return response, nil
}

func main() {

	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the gRPC service with the server
	pb.RegisterGRPCServiceServer(s, server{})	

	// Start accepting incoming requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
