package client

import (
	"context"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	client proto.StatsServiceClient
}

func NewGRPCClient(addr string) (*GRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCClient{client: proto.NewStatsServiceClient(conn)}, nil
}

func (c *GRPCClient) AddCall(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error) {
	return c.client.AddCall(ctx, req)
}

func (c *GRPCClient) GetStats(ctx context.Context, req *proto.StatsRequest) (*proto.StatsResponse, error) {
	return c.client.GetStats(ctx, req)
}

func (c *GRPCClient) CreateService(ctx context.Context, req *proto.ServiceRequest) (*proto.ServiceResponse, error) {
	return c.client.CreateService(ctx, req)
}
