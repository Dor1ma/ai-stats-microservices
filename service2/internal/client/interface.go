package client

import (
	"context"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
)

type GRPCClient interface {
	AddCall(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error)
	GetStats(ctx context.Context, req *proto.StatsRequest) (*proto.StatsResponse, error)
	CreateService(ctx context.Context, req *proto.ServiceRequest) (*proto.ServiceResponse, error)
}
