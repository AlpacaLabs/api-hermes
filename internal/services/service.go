package services

import (
	"github.com/AlpacaLabs/hermes/internal/config"
)

type Service struct {
	config config.Config
}

func NewService(config config.Config) Service {
	return Service{
		config: config,
	}
}
