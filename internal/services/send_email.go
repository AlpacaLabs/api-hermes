package services

import (
	"context"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) SendEmail(ctx context.Context, request hermesV1.SendEmailRequest) (*hermesV1.SendEmailResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
