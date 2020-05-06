package app

import (
	"sync"

	"github.com/AlpacaLabs/api-hermes/internal/grpc"

	"github.com/AlpacaLabs/api-hermes/internal/configuration"
	"github.com/AlpacaLabs/api-hermes/internal/http"
	"github.com/AlpacaLabs/api-hermes/internal/service"
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
	svc := service.NewService(a.config)

	var wg sync.WaitGroup

	wg.Add(1)
	httpServer := http.NewServer(a.config, svc)
	httpServer.Run()

	wg.Add(1)
	grpcServer := grpc.NewServer(a.config, svc)
	grpcServer.Run()

	wg.Wait()
}
