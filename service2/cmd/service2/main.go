package main

import (
	"fmt"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/client"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/config"
	"github.com/Dor1ma/ai-stats-microservices/service2/internal/handler"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	grpcAddress := fmt.Sprintf("%s:%s", cfg.GRPCHost, cfg.GRPCPort)

	grpcClient, err := client.NewGRPCClient(grpcAddress)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	handlers := handler.NewHandlers(grpcClient)

	http.HandleFunc("/call", handlers.AddCall)
	http.HandleFunc("/calls", handlers.GetCalls)
	http.HandleFunc("/service", handlers.CreateService)

	log.Printf("HTTP server is running on :%s\n", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
