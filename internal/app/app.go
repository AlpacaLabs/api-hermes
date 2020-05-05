package app

import (
	"sync"

	"github.com/AlpacaLabs/hermes/internal/grpc"

	"github.com/AlpacaLabs/hermes/internal/configuration"
	"github.com/AlpacaLabs/hermes/internal/http"
	"github.com/AlpacaLabs/hermes/internal/services"
)

type App struct {
	config configuration.Config
}

func NewApp(c configuration.Config) App {
	return App{
		config: c,
	}
}

func (a App) Run() {
	svc := services.NewService(a.config)

	var wg sync.WaitGroup

	wg.Add(1)
	httpServer := http.NewServer(a.config, svc)
	httpServer.Run()

	wg.Add(1)
	grpcServer := grpc.NewServer(a.config, svc)
	grpcServer.Run()

	wg.Wait()
}
