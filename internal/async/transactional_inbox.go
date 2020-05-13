package async

import (
	"context"
	"fmt"

	"github.com/AlpacaLabs/api-hermes/internal/configuration"
	"github.com/AlpacaLabs/api-hermes/internal/db"
	goKafka "github.com/AlpacaLabs/go-kafka"
	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	log "github.com/sirupsen/logrus"
)

func HandleSendEmailRequests(config configuration.Config) {
	ctx := context.TODO()

	groupID := config.AppName
	topic := "send-email-request"
	brokers := []string{
		fmt.Sprintf("%s:%d", config.KafkaConfig.Host, config.KafkaConfig.Port),
	}

	dbConn, err := config.SQLConfig.Connect()
	if err != nil {
		log.Fatalf("failed to connect to SQL: %v", err)
	}
	dbClient := db.NewClient(dbConn)

	err = goKafka.ProcessKafkaMessages(ctx, goKafka.ProcessKafkaMessagesInput{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		ProcessFunc: handleSendEmailRequest(dbClient),
	})
	if err != nil {
		log.Errorf("%v", err)
	}
}

func HandleSendSmsRequests(config configuration.Config) {
	ctx := context.TODO()

	groupID := config.AppName
	topic := "send-sms-request"
	brokers := []string{
		fmt.Sprintf("%s:%d", config.KafkaConfig.Host, config.KafkaConfig.Port),
	}

	dbConn, err := config.SQLConfig.Connect()
	if err != nil {
		log.Fatalf("failed to connect to SQL: %v", err)
	}
	dbClient := db.NewClient(dbConn)

	err = goKafka.ProcessKafkaMessages(ctx, goKafka.ProcessKafkaMessagesInput{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		ProcessFunc: handleSendSmsRequest(dbClient),
	})
	if err != nil {
		log.Errorf("%v", err)
	}
}

func handleSendEmailRequest(dbClient db.Client) func(context.Context, goKafka.Message) {
	return func(ctx context.Context, message goKafka.Message) {
		// Convert kafka.Message to Protocol Buffer
		pb := hermesV1.SendEmailRequest{}
		if err := message.Unmarshal(&pb); err != nil {
			log.Errorf("failed to unmarshal protobuf from kafka message: %v", err)
		}

		err := dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
			return tx.SaveSendEmailRequest(ctx, pb)
		})
		if err != nil {
			log.Errorf("failed to process kafka message in transaction: %v", err)
		}
	}
}

func handleSendSmsRequest(dbClient db.Client) func(context.Context, goKafka.Message) {
	return func(ctx context.Context, message goKafka.Message) {
		// Convert kafka.Message to Protocol Buffer
		pb := hermesV1.SendSmsRequest{}
		if err := message.Unmarshal(&pb); err != nil {
			log.Errorf("failed to unmarshal protobuf from kafka message: %v", err)
		}

		err := dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
			return tx.SaveSendSmsRequest(ctx, pb)
		})
		if err != nil {
			log.Errorf("failed to process kafka message in transaction: %v", err)
		}
	}
}
