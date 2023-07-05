package main

import (
    "context"
    "log"
    "google.golang.org/grpc"
)

// Replace "address:port" with the actual address and port of your gRPC server.
conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
if err != nil {
    log.Fatalf("Failed to connect: %v", err)
}
defer conn.Close()

// Create a new instance of the gRPC service client.
client := pb.NewGRPCServiceClient(conn)

// Create a request object with the necessary data.
request := &pb.Request{
    // Set the fields of the request object.
}

// Call the desired gRPC method on the client.
response, err := client.SendFeedback(context.Background(), request)
if err != nil {
    log.Fatalf("Failed to send feedback: %v", err)
}

// Handle the response received from the server.
// Use the fields of the response object as needed.
fmt.Printf("Response: Status=%v, Message=%v\n", response.Status, response.Message)


