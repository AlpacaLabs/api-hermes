package service

import (
	"github.com/AlpacaLabs/api-hermes/internal/configuration"
)

type Service struct {
	config configuration.Config
}

func NewService(config configuration.Config) Service {
	return Service{
		config: config,
	}
}
