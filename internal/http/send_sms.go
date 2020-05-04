package http

import (
	"net/http"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	"github.com/golang/protobuf/jsonpb"
)

func (s Server) SendSms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestBody hermesV1.SendSmsRequest
	if err := jsonpb.Unmarshal(r.Body, &requestBody); err != nil {
		// TODO return error
	}

	_, err := s.service.SendSms(ctx, requestBody)
	if err != nil {
		// TODO return error
	}

	// TODO return response
}
