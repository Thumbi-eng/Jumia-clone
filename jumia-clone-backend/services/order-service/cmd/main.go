package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "path/to/your/proto/order" // Update with the correct import path for your order.proto
)

func main() {
    // Set up a listener on the specified port
    lis, err := net.Listen("tcp", ":50051") // Change the port as needed
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Create a new gRPC server
    s := grpc.NewServer()

    // Register your service handlers here
    // pb.RegisterOrderServiceServer(s, &yourOrderServiceImplementation{}) // Uncomment and implement your service

    log.Println("Order service is running on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}