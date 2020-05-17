package async

import (
	"context"

	"github.com/AlpacaLabs/api-hermes/pkg/topic"

	"github.com/AlpacaLabs/api-hermes/internal/service"

	"github.com/AlpacaLabs/api-hermes/internal/configuration"
	goKafka "github.com/AlpacaLabs/go-kafka"
	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	log "github.com/sirupsen/logrus"
)

func HandleSendEmailRequests(config configuration.Config, s service.Service) {
	readFromTopic(topic.TopicForSendEmailRequest, config, handleSendEmailRequest(s))
}

func HandleSendSmsRequests(config configuration.Config, s service.Service) {
	readFromTopic(topic.TopicForSendSmsRequest, config, handleSendSmsRequest(s))
}

func handleSendEmailRequest(s service.Service) func(context.Context, goKafka.Message) {
	return func(ctx context.Context, message goKafka.Message) {
		// Convert kafka.Message to Protocol Buffer
		pb := hermesV1.SendEmailRequest{}
		if err := message.Unmarshal(&pb); err != nil {
			log.Errorf("failed to unmarshal protobuf from kafka message: %v", err)
		}

		if _, err := s.SendEmail(ctx, pb); err != nil {
			log.Errorf("failed to process kafka message in transaction: %v", err)
		}
	}
}

func handleSendSmsRequest(s service.Service) func(context.Context, goKafka.Message) {
	return func(ctx context.Context, message goKafka.Message) {
		// Convert kafka.Message to Protocol Buffer
		pb := hermesV1.SendSmsRequest{}
		if err := message.Unmarshal(&pb); err != nil {
			log.Errorf("failed to unmarshal protobuf from kafka message: %v", err)
		}

		if _, err := s.SendSms(ctx, pb); err != nil {
			log.Errorf("failed to process kafka message in transaction: %v", err)
		}
	}
}
