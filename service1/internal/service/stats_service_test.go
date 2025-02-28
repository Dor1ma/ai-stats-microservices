package service

import (
	"context"
	"database/sql"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRepository) GetConnection() *sql.DB {
	args := m.Called()
	return args.Get(0).(*sql.DB)
}

func (m *MockRepository) AddCall(ctx context.Context, userID, serviceID int64) error {
	args := m.Called(ctx, userID, serviceID)
	return args.Error(0)
}

func (m *MockRepository) GetStats(ctx context.Context, userID, serviceID int64, page, limit int32) ([]models.Stat, error) {
	args := m.Called(ctx, userID, serviceID, page, limit)
	return args.Get(0).([]models.Stat), args.Error(1)
}

func (m *MockRepository) CreateService(ctx context.Context, name, description string) (int64, error) {
	args := m.Called(ctx, name, description)
	return args.Get(0).(int64), args.Error(1)
}

type StatsServiceTestSuite struct {
	suite.Suite
	repo    *MockRepository
	service StatsService
}

func (suite *StatsServiceTestSuite) SetupTest() {
	suite.repo = new(MockRepository)
	suite.service = NewStatsService(suite.repo)
}

func (suite *StatsServiceTestSuite) TestAddCall() {
	ctx := context.Background()
	userID := int64(1)
	serviceID := int64(2)

	suite.repo.On("AddCall", ctx, userID, serviceID).Return(nil)

	err := suite.service.AddCall(ctx, userID, serviceID)

	assert.NoError(suite.T(), err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *StatsServiceTestSuite) TestGetStats() {
	ctx := context.Background()
	userID := int64(1)
	serviceID := int64(2)
	page := int32(1)
	limit := int32(10)

	mockStats := []models.Stat{
		{UserID: userID, ServiceID: serviceID, Count: 5, ServiceName: "Service1", Price: 10},
		{UserID: userID, ServiceID: serviceID, Count: 3, ServiceName: "Service2", Price: 20},
	}

	suite.repo.On("GetStats", ctx, userID, serviceID, page, limit).Return(mockStats, nil)

	stats, total, err := suite.service.GetStats(ctx, userID, serviceID, page, limit)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), int64(110), total)
	assert.Len(suite.T(), stats, 2)
	assert.Equal(suite.T(), "Service1", stats[0].ServiceName)
	assert.Equal(suite.T(), int64(50), stats[0].TotalAmount)
	assert.Equal(suite.T(), "Service2", stats[1].ServiceName)
	assert.Equal(suite.T(), int64(60), stats[1].TotalAmount)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *StatsServiceTestSuite) TestCreateService() {
	ctx := context.Background()
	name := "TestService"
	description := "TestDescription"
	serviceID := int64(1)

	suite.repo.On("CreateService", ctx, name, description).Return(serviceID, nil)

	id, err := suite.service.CreateService(ctx, name, description)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), serviceID, id)
	suite.repo.AssertExpectations(suite.T())
}

func TestStatsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(StatsServiceTestSuite))
}
