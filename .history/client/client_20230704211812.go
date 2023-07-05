package main

import (
    "context"
    "log"
    "google.golang.org/grpc"
)

// Replace "address:port" with the actual address and port of your gRPC server.
conn, err := grpc.Dial("address:port", grpc.WithInsecure())
if err != nil {
    log.Fatalf("Failed to connect: %v", err)
}
defer conn.Close()

// Create a new instance of the gRPC service client.
client := pb.NewGRPCServiceClient(conn)



