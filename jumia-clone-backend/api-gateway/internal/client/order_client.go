package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderServiceClient struct {
	Conn *grpc.ClientConn
}

func NewOrderServiceClient(addr string) (*OrderServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &OrderServiceClient{
		Conn: conn,
	}, nil
}

func (c *OrderServiceClient) Close() error {
	return c.Conn.Close()
}
