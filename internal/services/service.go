package services

import (
	"github.com/AlpacaLabs/hermes/internal/configuration"
)

type Service struct {
	config configuration.Config
}

func NewService(config configuration.Config) Service {
	return Service{
		config: config,
	}
}
