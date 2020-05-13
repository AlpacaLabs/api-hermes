package async

import (
	"context"
	"fmt"

	"github.com/AlpacaLabs/api-hermes/internal/service"

	"github.com/AlpacaLabs/api-hermes/internal/configuration"
	goKafka "github.com/AlpacaLabs/go-kafka"
	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	log "github.com/sirupsen/logrus"
)

func HandleSendEmailRequests(config configuration.Config, s service.Service) {
	ctx := context.TODO()

	groupID := config.AppName
	topic := "send-email-request"
	brokers := []string{
		fmt.Sprintf("%s:%d", config.KafkaConfig.Host, config.KafkaConfig.Port),
	}

	err := goKafka.ProcessKafkaMessages(ctx, goKafka.ProcessKafkaMessagesInput{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		ProcessFunc: handleSendEmailRequest(s),
	})
	if err != nil {
		log.Errorf("%v", err)
	}
}

func HandleSendSmsRequests(config configuration.Config, s service.Service) {
	ctx := context.TODO()

	groupID := config.AppName
	topic := "send-sms-request"
	brokers := []string{
		fmt.Sprintf("%s:%d", config.KafkaConfig.Host, config.KafkaConfig.Port),
	}

	err := goKafka.ProcessKafkaMessages(ctx, goKafka.ProcessKafkaMessagesInput{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		ProcessFunc: handleSendSmsRequest(s),
	})
	if err != nil {
		log.Errorf("%v", err)
	}
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
