package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Conn *grpc.ClientConn
}

func NewProductServiceClient(addr string) (*ProductServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ProductServiceClient{
		Conn: conn,
	}, nil
}

func (c *ProductServiceClient) Close() error {
	return c.Conn.Close()
}
