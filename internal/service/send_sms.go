package service

import (
	"context"

	"github.com/sfreiberg/gotwilio"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	log "github.com/sirupsen/logrus"
)

func (s Service) SendSms(ctx context.Context, request hermesV1.SendSmsRequest) (*hermesV1.SendSmsResponse, error) {
	if !s.config.SMSEnabled {
		log.Infof("SMS sending is disabled. Sending '%s' a message: '%s'", request.To, request.Message)
		return &hermesV1.SendSmsResponse{}, nil
	}

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
