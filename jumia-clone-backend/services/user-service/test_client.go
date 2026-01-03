package main

import (
	"context"
	"log"
	"time"

	pb "jumia-clone-backend/services/user-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Testing User Service...")
	log.Println("======================")

	// Test 1: Register
	log.Println("\n1. Testing Registration...")
	registerResp, err := client.Register(ctx, &pb.RegisterRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Phone:     "+1234567890",
	})
	if err != nil {
		log.Fatalf("Registration failed: %v", err)
	}

	log.Printf("✓ Registration successful! User ID: %s", registerResp.UserId)
	userID := registerResp.UserId

	// Test 2: Login
	log.Println("\n2. Testing Login...")
	loginResp, err := client.Login(ctx, &pb.LoginRequest{
		Email:    "john.doe@example.com",
		Password: "password123",
	})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	log.Printf("✓ Login successful! Token: %s...", loginResp.Token[:30])

	// Test 3: Get User
	log.Println("\n3. Testing Get User...")
	getUserResp, err := client.GetUser(ctx, &pb.GetUserRequest{UserId: userID})
	if err != nil {
		log.Fatalf("Get user failed: %v", err)
	}

	log.Printf("✓ Get user successful! Name: %s %s", getUserResp.User.FirstName, getUserResp.User.LastName)

	log.Println("\n======================")
	log.Println("All tests completed! ✓")
}
