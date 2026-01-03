package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// UserServiceClient holds the gRPC connection to user service
type UserServiceClient struct {
	Conn *grpc.ClientConn
}

// NewUserServiceClient creates a new user service client
func NewUserServiceClient(address string) (*UserServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to user service: %v", err)
		return nil, err
	}

	log.Printf("Connected to user service at %s", address)
	return &UserServiceClient{Conn: conn}, nil
}

// Close closes the gRPC connection
func (c *UserServiceClient) Close() error {
	return c.Conn.Close()
}
