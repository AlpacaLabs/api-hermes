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

// TODO this model for streaming is a bit bottlenecked:
// for example, a pod can only process one message at a time

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
		Brokers: brokers,
		GroupID: groupID,
		Topic:   topic,
	}, handleSendEmailRequest(dbClient))
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
		Brokers: brokers,
		GroupID: groupID,
		Topic:   topic,
	}, handleSendSmsRequest(dbClient))
	if err != nil {
		log.Errorf("%v", err)
	}
}

func handleSendEmailRequest(dbClient db.Client) func(context.Context, goKafka.Message) error {
	return func(ctx context.Context, message goKafka.Message) error {
		// Convert kafka.Message to Protocol Buffer
		pb := &hermesV1.SendEmailRequest{}
		if err := message.Unmarshal(pb); err != nil {
			return err
		}

		return dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
			return tx.SaveSendEmailRequest(ctx, *pb)
		})
	}
}

func handleSendSmsRequest(dbClient db.Client) func(context.Context, goKafka.Message) error {
	return func(ctx context.Context, message goKafka.Message) error {
		// Convert kafka.Message to Protocol Buffer
		pb := &hermesV1.SendSmsRequest{}
		if err := message.Unmarshal(pb); err != nil {
			return err
		}

		return dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
			return tx.SaveSendSmsRequest(ctx, *pb)
		})
	}
}
