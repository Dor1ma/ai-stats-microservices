package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGRPCClient struct {
	mock.Mock
}

func (m *MockGRPCClient) CreateService(ctx context.Context, req *proto.ServiceRequest) (*proto.ServiceResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*proto.ServiceResponse), args.Error(1)
}

func (m *MockGRPCClient) AddCall(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*proto.CallResponse), args.Error(1)
}

func (m *MockGRPCClient) GetStats(ctx context.Context, req *proto.StatsRequest) (*proto.StatsResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*proto.StatsResponse), args.Error(1)
}

func TestHandlers_AddCall(t *testing.T) {
	logger.InitLogger()
	defer logger.Log.Sync()

	mockClient := new(MockGRPCClient)
	handler := NewHandlers(mockClient)

	reqBody := map[string]int64{
		"user_id":    1,
		"service_id": 2,
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	expectedResponse := &proto.CallResponse{Success: true}
	mockClient.On("AddCall", mock.Anything, &proto.CallRequest{UserId: 1, ServiceId: 2}).Return(expectedResponse, nil)

	req, err := http.NewRequest("POST", "/call", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.AddCall(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp proto.CallResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.True(t, resp.Success)

	mockClient.AssertExpectations(t)
}

func TestHandlers_GetCalls(t *testing.T) {
	logger.InitLogger()
	defer logger.Log.Sync()

	mockClient := new(MockGRPCClient)
	handler := NewHandlers(mockClient)

	userID := int64(1)
	serviceID := int64(2)
	page := 1
	limit := 10

	expectedResponse := &proto.StatsResponse{
		Stats: []*proto.Stat{
			{UserId: 1, ServiceId: 2, Count: 5, ServiceName: "Service1", TotalAmount: 50},
		},
		Total: 1,
	}
	mockClient.On("GetStats", mock.Anything, &proto.StatsRequest{UserId: &userID, ServiceId: &serviceID, Page: int32(page), Limit: int32(limit)}).Return(expectedResponse, nil)

	req, err := http.NewRequest("GET", "/calls?user_id=1&service_id=2&page=1&limit=10", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.GetCalls(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp proto.StatsResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)

	t.Logf("Expected Stats: %+v", expectedResponse.Stats)
	t.Logf("Actual Stats: %+v", resp.Stats)

	expectedStats := make([]proto.Stat, len(expectedResponse.Stats))
	actualStats := make([]proto.Stat, len(resp.Stats))

	for i := range expectedResponse.Stats {
		expectedStats[i] = *expectedResponse.Stats[i]
	}
	for i := range resp.Stats {
		actualStats[i] = *resp.Stats[i]
	}

	assert.ElementsMatch(t, expectedStats, actualStats)

	assert.EqualValues(t, expectedResponse.Total, resp.Total)

	mockClient.AssertExpectations(t)
}

func TestHandlers_CreateService(t *testing.T) {
	logger.InitLogger()
	defer logger.Log.Sync()

	mockClient := new(MockGRPCClient)
	handler := NewHandlers(mockClient)

	reqBody := map[string]interface{}{
		"name":        "TestService",
		"description": "TestDescription",
		"price":       100,
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	expectedResponse := &proto.ServiceResponse{Id: 1}
	mockClient.On("CreateService", mock.Anything, &proto.ServiceRequest{Name: "TestService", Description: "TestDescription", Price: 100}).Return(expectedResponse, nil)

	req, err := http.NewRequest("POST", "/service", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.CreateService(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var resp proto.ServiceResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Id, resp.Id)

	mockClient.AssertExpectations(t)
}
