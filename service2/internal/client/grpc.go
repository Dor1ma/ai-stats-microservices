package client

import (
	"context"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"google.golang.org/grpc"
)

type gRPCClient struct {
	client proto.StatsServiceClient
}

func NewGRPCClient(addr string) (GRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &gRPCClient{client: proto.NewStatsServiceClient(conn)}, nil
}

func (c *gRPCClient) AddCall(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error) {
	return c.client.AddCall(ctx, req)
}

func (c *gRPCClient) GetStats(ctx context.Context, req *proto.StatsRequest) (*proto.StatsResponse, error) {
	return c.client.GetStats(ctx, req)
}

func (c *gRPCClient) CreateService(ctx context.Context, req *proto.ServiceRequest) (*proto.ServiceResponse, error) {
	return c.client.CreateService(ctx, req)
}
