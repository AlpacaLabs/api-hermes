package services

import (
	"context"

	"github.com/sfreiberg/gotwilio"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
)

func (s Service) SendSms(ctx context.Context, request hermesV1.SendSmsRequest) (*hermesV1.SendSmsResponse, error) {
	twilioClient := gotwilio.NewTwilioClient(s.config.TwilioAccountSID, s.config.TwilioAuthToken)

	_, exception, err := twilioClient.SendSMS(s.config.TwilioPhoneNumber, request.To, request.Message, "", "")
	if err != nil {
		return nil, err
	}
	if exception != nil {
		return nil, exception
	}

	return &hermesV1.SendSmsResponse{}, nil
}
