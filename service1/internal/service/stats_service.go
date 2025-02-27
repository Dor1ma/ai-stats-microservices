package service

import (
	"context"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/repository"
)

type statsService struct {
	repo repository.Repository
}

func NewStatsService(repo repository.Repository) StatsService {
	return &statsService{repo: repo}
}

func (s *statsService) AddCall(ctx context.Context, userID, serviceID int64) error {
	return s.repo.AddCall(ctx, userID, serviceID)
}

func (s *statsService) GetStats(ctx context.Context, userID, serviceID int64, page, limit int32) ([]*proto.Stat, int64, error) {
	stats, err := s.repo.GetStats(ctx, userID, serviceID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	var protoStats []*proto.Stat
	for _, stat := range stats {
		protoStats = append(protoStats, &proto.Stat{
			UserId:      stat.UserID,
			ServiceId:   stat.ServiceID,
			Count:       stat.Count,
			ServiceName: stat.ServiceName,
			TotalAmount: stat.Count * stat.Price,
		})
		total += stat.Count * stat.Price
	}

	return protoStats, total, nil
}

func (s *statsService) CreateService(ctx context.Context, name, description string) (int64, error) {
	return s.repo.CreateService(ctx, name, description)
}
