package handler

import (
	"context"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"

	"github.com/Dor1ma/ai-stats-microservices/service1/internal/service"
)

type StatsHandler struct {
	service service.StatsService
	proto.UnimplementedStatsServiceServer
}

func NewStatsHandler(service service.StatsService) *StatsHandler {
	return &StatsHandler{service: service}
}

func (s *StatsHandler) AddCall(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error) {
	err := s.service.AddCall(ctx, req.UserId, req.ServiceId)
	return &proto.CallResponse{Success: err == nil}, err
}

func (s *StatsHandler) GetStats(ctx context.Context, req *proto.StatsRequest) (*proto.StatsResponse, error) {
	stats, total, err := s.service.GetStats(ctx, *req.UserId, *req.ServiceId, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	return &proto.StatsResponse{Stats: stats, Total: total}, nil
}

func (s *StatsHandler) CreateService(ctx context.Context, req *proto.ServiceRequest) (*proto.ServiceResponse, error) {
	id, err := s.service.CreateService(ctx, req.Name, req.Description)
	return &proto.ServiceResponse{Id: id}, err
}
