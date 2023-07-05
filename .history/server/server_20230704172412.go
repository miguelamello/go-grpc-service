package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	ppb "go_grpc_service/go_grpc_service"
)

type server struct{}

func (s *server) GRPCService(ping *ppb.Request, stream ppb.GRPCServiceServer) error {
	for {
		// Receive the client's ping message
		fmt.Printf("Received ping: %s\n", ping.Message)

		// Construct the pong message
		pong := &ppb.{
			Message: "Pong: " + ping.Message,
		}

		// Send the pong message back to the client
		if err := stream.Send(pong); err != nil {
			return err
		}

		// Receive the next ping message from the client
		ping, err := stream.Recv()
		if err != nil {
			return err
		}
	}
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the PingPong service with the server
	ppb.RegisterGRPCServiceServer(s, &server{})

	// Start accepting incoming requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
