package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	HTTPPort string
	GRPCHost string
	GRPCPort string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	grpcHost := os.Getenv("GRPC_HOST")
	if grpcHost == "" {
		grpcHost = "localhost"
		log.Println("Warning: GRPC_HOST environment variable not set, using default localhost")
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
		log.Println("Info: GRPC_PORT is not set, using default 50051")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
		log.Println("Info: HTTP_PORT is not set, using default 8080")
	}

	return &Config{
		GRPCHost: grpcHost,
		GRPCPort: grpcPort,
		HTTPPort: httpPort,
	}
}
