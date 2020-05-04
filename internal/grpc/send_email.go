package grpc

import (
	"context"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
)

func (s Server) SendEmail(ctx context.Context, request *hermesV1.SendEmailRequest) (*hermesV1.SendEmailResponse, error) {
	return s.service.SendEmail(ctx, *request)
}
