package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CartServiceClient struct {
	Conn *grpc.ClientConn
}

func NewCartServiceClient(addr string) (*CartServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &CartServiceClient{
		Conn: conn,
	}, nil
}

func (c *CartServiceClient) Close() error {
	return c.Conn.Close()
}
