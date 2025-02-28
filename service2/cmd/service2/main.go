package main

import (
	"fmt"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/client"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/config"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/handler"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/logger"
	"net/http"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Starting service2...")
	cfg := config.LoadConfig()

	grpcAddress := fmt.Sprintf("%s:%s", cfg.GRPCHost, cfg.GRPCPort)

	logger.Log.Info(fmt.Sprintf("grpc address: %s", grpcAddress))

	grpcClient, err := client.NewGRPCClient(grpcAddress)
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Failed to connect to gRPC server: %v", err))
	}

	handlers := handler.NewHandlers(grpcClient)

	http.HandleFunc("/call", handlers.AddCall)
	http.HandleFunc("/calls", handlers.GetCalls)
	http.HandleFunc("/service", handlers.CreateService)

	logger.Log.Info(fmt.Sprintf("HTTP server is running on :%s\n", cfg.HTTPPort))
	if err := http.ListenAndServe(":"+cfg.HTTPPort, nil); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Failed to start HTTP server: %v", err))
	}
}
