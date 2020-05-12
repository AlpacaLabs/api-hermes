package async

import (
	"context"

	"github.com/AlpacaLabs/api-hermes/internal/configuration"
	"github.com/AlpacaLabs/api-hermes/internal/db"
	goKafka "github.com/AlpacaLabs/go-kafka"
	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// TODO this model for streaming is a bit bottlenecked:
// for example, a pod can only process one message at a time

func HandleSendEmailRequests(config configuration.Config) {
	ctx := context.TODO()

	groupID := config.AppName
	topic := "send-email-request"

	reader, closer := goKafka.GetReader(goKafka.GetReaderInput{
		Host:    config.KafkaConfig.Host,
		Port:    config.KafkaConfig.Port,
		GroupID: groupID,
		Topic:   topic,
	})
	defer closer()

	dbConn, err := config.SQLConfig.Connect()
	if err != nil {
		log.Fatalf("failed to connect to SQL: %v", err)
	}
	dbClient := db.NewClient(dbConn)

	err = goKafka.ProcessKafkaMessages(ctx, reader, handleSendEmailRequest(dbClient))
	if err != nil {
		log.Errorf("%v", err)
	}
}

func HandleSendSmsRequests(config configuration.Config) {
	ctx := context.TODO()

	groupID := config.AppName
	topic := "send-sms-request"

	reader, closer := goKafka.GetReader(goKafka.GetReaderInput{
		Host:    config.KafkaConfig.Host,
		Port:    config.KafkaConfig.Port,
		GroupID: groupID,
		Topic:   topic,
	})
	defer closer()

	dbConn, err := config.SQLConfig.Connect()
	if err != nil {
		log.Fatalf("failed to connect to SQL: %v", err)
	}
	dbClient := db.NewClient(dbConn)

	err = goKafka.ProcessKafkaMessages(ctx, reader, handleSendSmsRequest(dbClient))
	if err != nil {
		log.Errorf("%v", err)
	}
}

func handleSendEmailRequest(dbClient db.Client) func(context.Context, kafka.Message) error {
	return func(ctx context.Context, message kafka.Message) error {
		// Convert kafka.Message to Protocol Buffer
		pb := &hermesV1.SendEmailRequest{}
		if err := proto.Unmarshal(message.Value, pb); err != nil {
			return err
		}

		return dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
			return tx.SaveSendEmailRequest(ctx, *pb)
		})
	}
}

func handleSendSmsRequest(dbClient db.Client) func(context.Context, kafka.Message) error {
	return func(ctx context.Context, message kafka.Message) error {
		// Convert kafka.Message to Protocol Buffer
		pb := &hermesV1.SendSmsRequest{}
		if err := proto.Unmarshal(message.Value, pb); err != nil {
			return err
		}

		return dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
			return tx.SaveSendSmsRequest(ctx, *pb)
		})
	}
}
