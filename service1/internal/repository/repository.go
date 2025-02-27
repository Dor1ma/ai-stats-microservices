package repository

import (
	"context"
	"database/sql"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/models"
)

type Repository interface {
	CreateService(ctx context.Context, name, description string) (int64, error)
	AddCall(ctx context.Context, userID, serviceID int64) error
	GetStats(ctx context.Context, userID, serviceID int64, offset, limit int32) ([]models.Stat, error)
	Close() error
	GetConnection() *sql.DB
}
