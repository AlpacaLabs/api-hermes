package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/AlpacaLabs/api-hermes/internal/app"
	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"

	"github.com/AlpacaLabs/api-hermes/internal/configuration"
	. "github.com/smartystreets/goconvey/convey"
)

var config configuration.Config

func TestMain(m *testing.M) {
	config = configuration.LoadConfig()

	a := app.NewApp(config)

	var wg sync.WaitGroup

	wg.Add(1)
	go a.Run()

	// TODO verify server is healthy before running tests

	code := m.Run()

	os.Exit(code)
}

func Test_KafkaTopicRead(t *testing.T) {
	Convey("Given a Kafka topic", t, func(c C) {
		ctx := context.TODO()

		topic := "send-email-request"
		brokers := []string{
			fmt.Sprintf("%s:%d", config.KafkaConfig.Host, config.KafkaConfig.Port),
		}

		w := kafka.NewWriter(kafka.WriterConfig{
			Brokers: brokers,
			Topic:   topic,
		})

		pb := &hermesV1.SendEmailRequest{
			Email: &hermesV1.Email{
				Body: &hermesV1.Body{
					Intros: []string{"hi"},
				},
			},
		}

		b, err := proto.Marshal(pb)
		So(err, ShouldBeNil)

		err = w.WriteMessages(ctx, kafka.Message{
			Value: b,
		})
		So(err, ShouldBeNil)
	})
}
