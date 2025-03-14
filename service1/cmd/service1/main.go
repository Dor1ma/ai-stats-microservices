package main

import (
	"fmt"
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/config"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/handler"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/logger"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/repository"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Starting service1...")
	cfg := config.LoadConfig()

	repo, err := repository.NewPostgresRepository(cfg.DBConnStr)
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	defer repo.Close()

	statsService := service.NewStatsService(repo)
	statsHandler := handler.NewStatsHandler(statsService)

	grpcServer := grpc.NewServer()
	proto.RegisterStatsServiceServer(grpcServer, statsHandler)

	address := "0.0.0.0:" + cfg.GRPCPort
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Failed to listen: %v", err))
	}

	logger.Log.Info(fmt.Sprintf("gRPC server is running on :%v", lis.Addr()))
	if err := grpcServer.Serve(lis); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Failed to serve: %v", err))
	}
}
