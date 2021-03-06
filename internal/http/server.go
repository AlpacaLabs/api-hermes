package http

import (
	"fmt"
	"net/http"

	"github.com/AlpacaLabs/api-hermes/internal/configuration"
	"github.com/AlpacaLabs/api-hermes/internal/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	config  configuration.Config
	service service.Service
}

func NewServer(config configuration.Config, service service.Service) Server {
	return Server{
		config:  config,
		service: service,
	}
}

func (s Server) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/send-email", s.SendEmail).Methods(http.MethodPost)
	r.HandleFunc("/send-sms", s.SendSms).Methods(http.MethodPost)

	addr := fmt.Sprintf(":%d", s.config.HTTPPort)
	log.Infof("Listening for HTTP on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
