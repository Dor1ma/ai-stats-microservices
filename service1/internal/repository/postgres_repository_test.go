package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/Dor1ma/ai-stats-microservices/service1/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestPostgresRepository_CreateService(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &PostgresRepository{conn: db}

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(`^INSERT INTO services \(name, description, price\) VALUES \(\$1, \$2, \$3\) RETURNING id$`).
			WithArgs("test-service", "test description", 100).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		id, err := repo.CreateService(context.Background(), "test-service", "test description", 100)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := fmt.Errorf("database error")
		mock.ExpectQuery(`^INSERT INTO services \(name, description, price\) VALUES \(\$1, \$2, \$3\) RETURNING id$`).
			WithArgs("fail-service", "", 100).
			WillReturnError(expectedErr)

		id, err := repo.CreateService(context.Background(), "fail-service", "", 100)

		assert.ErrorIs(t, err, expectedErr)
		assert.Equal(t, int64(0), id)
	})

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresRepository_AddCall(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &PostgresRepository{conn: db}

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(`^INSERT INTO stats \(user_id, service_id, count\) VALUES \(\$1, \$2, 1\) ON CONFLICT \(user_id, service_id\) DO UPDATE SET count = stats\.count \+ 1$`).
			WithArgs(1, 2).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.AddCall(context.Background(), 1, 2)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := fmt.Errorf("exec error")
		mock.ExpectExec(`^INSERT INTO stats \(user_id, service_id, count\) VALUES \(\$1, \$2, 1\) ON CONFLICT \(user_id, service_id\) DO UPDATE SET count = stats\.count \+ 1$`).
			WithArgs(3, 4).
			WillReturnError(expectedErr)

		err := repo.AddCall(context.Background(), 3, 4)

		assert.ErrorIs(t, err, expectedErr)
	})

	assert.NoError(t, mock.ExpectationsWereMet())
}
