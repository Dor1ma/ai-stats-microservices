package handler

import (
	"context"
	"encoding/json"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/client"
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
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.grpcClient.AddCall(context.Background(), &proto.CallRequest{
		UserId:    req.UserID,
		ServiceId: req.ServiceID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
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

	resp, err := h.grpcClient.GetStats(context.Background(), &proto.StatsRequest{
		UserId:    &userID,
		ServiceId: &serviceID,
		Page:      int32(page),
		Limit:     int32(limit),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handlers) CreateService(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.grpcClient.CreateService(context.Background(), &proto.ServiceRequest{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
