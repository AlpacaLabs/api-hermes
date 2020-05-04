package grpc

import (
	"fmt"
	"net"

	"github.com/AlpacaLabs/hermes/internal/config"
	"github.com/AlpacaLabs/hermes/internal/services"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config  config.Config
	service services.Service
}

func NewServer(config config.Config, service services.Service) Server {
	return Server{
		config:  config,
		service: service,
	}
}

func (s Server) Run() {
	address := fmt.Sprintf(":%d", s.config.GrpcPort)

	log.Printf("Listening on %s\n", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Starting gRPC server...")
	grpcServer := grpc.NewServer()

	// TODO Register our services

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Info("Registered gRPC services...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
