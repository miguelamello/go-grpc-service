package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "go_grpc_service/go_grpc_service"
)

func main() {

	// Replace "address:port" with the actual address and port of your gRPC server.
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new instance of the gRPC service client.
	client := pb.NewGRPCServiceClient(conn)

	// Create a request object with the necessary data.
	request := &pb.Request{
		// Set the fields of the request object.
		Message: "",
	}

	// Start a goroutine to send messages every 1 second.
	go func() {
		for {
			// Create a new UUID for each message.
			uuid := uuid.New()
			// Set the message field of the request object.
			request.Message = uuid.String()

			// Call the desired gRPC method on the client.
			response, err := client.SendFeedback(context.Background(), request)
			if err != nil {
				log.Fatalf("Failed to send feedback: %v", err)
			}

			// Handle the response received from the server.
			// Use the fields of the response object as needed.
			fmt.Println(response.Message)

			// Sleep for 1 second before sending the next message.
			time.Sleep(1 * time.Second)
		}
	}()

	// Wait indefinitely to keep the program running.
	select {}

}



