package service

import (
	"context"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
)

type StatsService interface {
	AddCall(ctx context.Context, userID, serviceID int64) error
	GetStats(ctx context.Context, userID, serviceID int64, page, limit int32) ([]*proto.Stat, int64, error)
	CreateService(ctx context.Context, name, description string) (int64, error)
}
