package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/models"
	_ "github.com/Dor1ma/ai-stats-microservices/service1/internal/models"
	_ "github.com/lib/pq"
)

const (
	createServiceQuery = `
		INSERT INTO services (name, description, price)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	addCallQuery = `
		INSERT INTO stats (user_id, service_id, count)
		VALUES ($1, $2, 1)
		ON CONFLICT (user_id, service_id)
		DO UPDATE SET count = stats.count + 1
	`

	getStatsQuery = `
		SELECT s.user_id, s.service_id, s.count, sv.name, sv.price
		FROM stats s
		JOIN services sv ON s.service_id = sv.id
		WHERE ($1 = 0 OR s.user_id = $1)
		AND ($2 = 0 OR s.service_id = $2)
		LIMIT $3 OFFSET $4
	`
)

type PostgresRepository struct {
	conn *sql.DB
}

func NewPostgresRepository(connString string) (Repository, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresRepository{conn: db}, nil
}

func (d *PostgresRepository) Close() error {
	return d.conn.Close()
}

func (d *PostgresRepository) GetConnection() *sql.DB {
	return d.conn
}

func (d *PostgresRepository) CreateService(ctx context.Context, name, description string, price int64) (int64, error) {
	var id int64
	err := d.conn.QueryRowContext(ctx, createServiceQuery, name, description, price).Scan(&id)
	return id, err
}

func (d *PostgresRepository) AddCall(ctx context.Context, userID, serviceID int64) error {
	_, err := d.conn.ExecContext(ctx, addCallQuery, userID, serviceID)
	return err
}

func (d *PostgresRepository) GetStats(ctx context.Context, userID, serviceID int64, offset, limit int32) ([]models.Stat, error) {
	if offset < 1 {
		offset = 0
	} else {
		offset -= 1
	}

	rows, err := d.conn.QueryContext(ctx, getStatsQuery, userID, serviceID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []models.Stat
	for rows.Next() {
		var stat models.Stat
		if err := rows.Scan(&stat.UserID, &stat.ServiceID, &stat.Count, &stat.ServiceName, &stat.Price); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}
	return stats, nil
}
