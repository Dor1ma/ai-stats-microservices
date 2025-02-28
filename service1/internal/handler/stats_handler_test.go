package handler

import (
	"context"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/logger"
	"testing"

	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	_ "github.com/Dor1ma/ai-stats-microservices/service1/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStatsService struct {
	mock.Mock
}

func (m *MockStatsService) AddCall(ctx context.Context, userID, serviceID int64) error {
	args := m.Called(ctx, userID, serviceID)
	return args.Error(0)
}

func (m *MockStatsService) GetStats(ctx context.Context, userID, serviceID int64, page, limit int32) ([]*proto.Stat, int64, error) {
	args := m.Called(ctx, userID, serviceID, page, limit)
	return args.Get(0).([]*proto.Stat), args.Get(1).(int64), args.Error(2)
}

func (m *MockStatsService) CreateService(ctx context.Context, name, description string, price int64) (int64, error) {
	args := m.Called(ctx, name, description, price)
	return args.Get(0).(int64), args.Error(1)
}

func TestAddCall(t *testing.T) {
	logger.InitLogger()
	defer logger.Log.Sync()

	mockService := new(MockStatsService)
	handler := NewStatsHandler(mockService)

	ctx := context.Background()
	req := &proto.CallRequest{
		UserId:    1,
		ServiceId: 2,
	}

	mockService.On("AddCall", ctx, req.UserId, req.ServiceId).Return(nil)

	resp, err := handler.AddCall(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockService.AssertExpectations(t)
}

func TestGetStats(t *testing.T) {
	logger.InitLogger()
	defer logger.Log.Sync()

	mockService := new(MockStatsService)
	handler := NewStatsHandler(mockService)

	userId := int64(1)
	serviceId := int64(2)

	ctx := context.Background()
	req := &proto.StatsRequest{
		UserId:    &userId,
		ServiceId: &serviceId,
		Page:      1,
		Limit:     10,
	}

	mockStats := []*proto.Stat{
		{UserId: 1, ServiceId: 2, Count: 5, ServiceName: "Service1", TotalAmount: 50},
		{UserId: 1, ServiceId: 2, Count: 3, ServiceName: "Service2", TotalAmount: 60},
	}
	var total int64 = 110

	mockService.On("GetStats", ctx, *req.UserId, *req.ServiceId, req.Page, req.Limit).Return(mockStats, total, nil)

	resp, err := handler.GetStats(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, mockStats, resp.Stats)
	assert.Equal(t, total, resp.Total)
	mockService.AssertExpectations(t)
}

func TestCreateService(t *testing.T) {
	logger.InitLogger()
	defer logger.Log.Sync()

	mockService := new(MockStatsService)
	handler := NewStatsHandler(mockService)

	ctx := context.Background()
	req := &proto.ServiceRequest{
		Name:        "TestService",
		Description: "TestDescription",
		Price:       100,
	}

	var serviceID int64 = 1
	mockService.On("CreateService", ctx, req.Name, req.Description, req.Price).Return(serviceID, nil)

	resp, err := handler.CreateService(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, serviceID, resp.Id)
	mockService.AssertExpectations(t)
}
