package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	ppb "path/to/generated/package/pingpong"
)

func main() {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new PingPong client
	client := ppb.NewPingPongClient(conn)

	// Send the initial ping message to the server
	stream, err := client.SendPing(context.Background())
	if err != nil {
		log.Fatalf("Failed to send ping: %v", err)
	}

	// Send and receive ping-pong messages
	for i := 0; i < 5; i++ {
		// Construct the ping message
		ping := &ppb.Ping{
			Message: "Ping",
		}

		// Send the ping message to the server
		if err := stream.Send(ping); err != nil {
			log.Fatalf("Failed to send ping: %v", err)
		}

		// Receive the pong message from the server
		pong, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive pong: %v", err)
		}

		// Print the received pong message
		log.Printf("Received pong: %s\n", pong.Message)

		// Sleep for a second before sending the next ping
		time.Sleep(time.Second)
	}

	// Close the stream
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Failed to close stream: %v", err)
	}
}
