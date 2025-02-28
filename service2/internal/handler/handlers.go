package handler

import (
	"context"
	"encoding/json"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/client"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handlers struct {
	grpcClient *client.GRPCClient
}

func NewHandlers(grpcClient *client.GRPCClient) *Handlers {
	return &Handlers{grpcClient: grpcClient}
}

func (h *Handlers) AddCall(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID    int64 `json:"user_id"`
		ServiceID int64 `json:"service_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	logger.Log.Info("AddCall request received", zap.Int64("user_id", req.UserID), zap.Int64("service_id", req.ServiceID))

	response, err := h.grpcClient.AddCall(context.Background(), &proto.CallRequest{
		UserId:    req.UserID,
		ServiceId: req.ServiceID,
	})
	if err != nil {
		logger.Log.Error("Failed to add call", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.Info("Call added successfully", zap.Any("response", response))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Log.Error("Failed to encode response", zap.Error(err))
	}
}

func (h *Handlers) GetCalls(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	serviceID, _ := strconv.ParseInt(r.URL.Query().Get("service_id"), 10, 64)
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	logger.Log.Info("GetCalls request received", zap.Int64("user_id", userID), zap.Int64("service_id", serviceID), zap.Int("page", page), zap.Int("limit", limit))

	resp, err := h.grpcClient.GetStats(context.Background(), &proto.StatsRequest{
		UserId:    &userID,
		ServiceId: &serviceID,
		Page:      int32(page),
		Limit:     int32(limit),
	})
	if err != nil {
		logger.Log.Error("Failed to get calls", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.Info("Calls retrieved successfully", zap.Any("response", resp))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Log.Error("Failed to encode response", zap.Error(err))
	}
}

func (h *Handlers) CreateService(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	logger.Log.Info("CreateService request received", zap.String("name", req.Name), zap.String("description", req.Description))

	resp, err := h.grpcClient.CreateService(context.Background(), &proto.ServiceRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	})
	if err != nil {
		logger.Log.Error("Failed to create service", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.Info("Service created successfully", zap.Any("response", resp))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Log.Error("Failed to encode response", zap.Error(err))
	}
}
