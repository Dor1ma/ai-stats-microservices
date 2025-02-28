package main

import (
	proto "github.com/Dor1ma/ai-stats-microservices/proto/gen"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/config"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/handler"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/repository"
	"github.com/Dor1ma/ai-stats-microservices/service1/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.LoadConfig()

	repo, err := repository.NewPostgresRepository(cfg.DBConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer repo.Close()

	statsService := service.NewStatsService(repo)
	statsHandler := handler.NewStatsHandler(statsService)

	grpcServer := grpc.NewServer()
	proto.RegisterStatsServiceServer(grpcServer, statsHandler)

	address := "0.0.0.0:" + cfg.GRPCPort
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("gRPC server is running on :%v\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
