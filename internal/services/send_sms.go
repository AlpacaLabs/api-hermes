package services

import (
	"context"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) SendSms(ctx context.Context, request hermesV1.SendSmsRequest) (*hermesV1.SendSmsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
