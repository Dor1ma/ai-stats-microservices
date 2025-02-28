package handler

import (
	"context"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/logger"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/service"
	"go.uber.org/zap"
)

type StatsHandler struct {
	service service.StatsService
	proto.UnimplementedStatsServiceServer
}

func NewStatsHandler(service service.StatsService) *StatsHandler {
	return &StatsHandler{service: service}
}

func (s *StatsHandler) AddCall(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error) {
	logger.Log.Info("AddCall request received", zap.Int64("user_id", req.UserId), zap.Int64("service_id", req.ServiceId))

	err := s.service.AddCall(ctx, req.UserId, req.ServiceId)
	if err != nil {
		logger.Log.Error("Failed to add call", zap.Int64("user_id", req.UserId), zap.Int64("service_id", req.ServiceId), zap.Error(err))
		return &proto.CallResponse{Success: false}, err
	}

	logger.Log.Info("Call added successfully", zap.Int64("user_id", req.UserId), zap.Int64("service_id", req.ServiceId))
	return &proto.CallResponse{Success: true}, nil
}

func (s *StatsHandler) GetStats(ctx context.Context, req *proto.StatsRequest) (*proto.StatsResponse, error) {
	logger.Log.Info("GetStats request received",
		zap.Int64p("user_id", req.UserId),
		zap.Int64p("service_id", req.ServiceId),
		zap.Int32("page", req.Page),
		zap.Int32("limit", req.Limit),
	)

	stats, total, err := s.service.GetStats(ctx, *req.UserId, *req.ServiceId, req.Page, req.Limit)
	if err != nil {
		logger.Log.Error("Failed to get stats",
			zap.Int64p("user_id", req.UserId),
			zap.Int64p("service_id", req.ServiceId),
			zap.Error(err),
		)
		return nil, err
	}

	logger.Log.Info("Stats retrieved successfully",
		zap.Int64p("user_id", req.UserId),
		zap.Int64p("service_id", req.ServiceId),
		zap.Int32("total", int32(total)),
	)
	return &proto.StatsResponse{Stats: stats, Total: total}, nil
}

func (s *StatsHandler) CreateService(ctx context.Context, req *proto.ServiceRequest) (*proto.ServiceResponse, error) {
	logger.Log.Info("CreateService request received",
		zap.String("name", req.Name),
		zap.String("description", req.Description),
	)

	id, err := s.service.CreateService(ctx, req.Name, req.Description, req.Price)
	if err != nil {
		logger.Log.Error("Failed to create service",
			zap.String("name", req.Name),
			zap.String("description", req.Description),
			zap.Error(err),
		)
		return &proto.ServiceResponse{Id: id}, err
	}

	logger.Log.Info("Service created successfully",
		zap.String("name", req.Name),
		zap.String("description", req.Description),
		zap.Int64("id", id),
	)
	return &proto.ServiceResponse{Id: id}, nil
}
