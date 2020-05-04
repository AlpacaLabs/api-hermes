package grpc

import (
	"context"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
)

func (s Server) SendSms(ctx context.Context, request *hermesV1.SendSmsRequest) (*hermesV1.SendSmsResponse, error) {
	return s.service.SendSms(ctx, *request)
}
