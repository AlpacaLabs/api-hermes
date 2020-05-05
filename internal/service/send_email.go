package service

import (
	"context"

	"github.com/golang/protobuf/jsonpb"
	log "github.com/sirupsen/logrus"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) SendEmail(ctx context.Context, request hermesV1.SendEmailRequest) (*hermesV1.SendEmailResponse, error) {
	if !s.config.EmailEnabled {
		primaryRecipient := request.Email.To[0]

		m := jsonpb.Marshaler{}
		s, err := m.MarshalToString(request.Email)
		if err != nil {
			return nil, err
		}

		log.Infof("Email sending is disabled. Sending '%s' a message: '%s'", primaryRecipient, s)
		return &hermesV1.SendEmailResponse{}, nil
	}

	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
