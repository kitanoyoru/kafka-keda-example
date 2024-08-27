package consumer

import (
	"log/slog"

	"github.com/kitanoyoru/kafka-keda-example/pkg/cmdutils"
	"github.com/kitanoyoru/kafka-keda-example/pkg/config"
	"github.com/kitanoyoru/kafka-keda-example/pkg/gen/go/user"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hamba/avro"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cfg = config.Consumer{
	SchemaPaths:      []string{"/usr/local/lib/schema/avro/user.avsc"},
	BootstrapServers: "kafka:29092",
	Topic:            "kita.user",
	ConsumerGroupID:  "user-consumer-group",
}

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "consumer",
		RunE: func(cmd *cobra.Command, args []string) error {
			schema, err := avro.ParseFiles(cfg.SchemaPaths...)
			if err != nil {
				return errors.Wrap(err, "failed to parse Avro schema")
			}

			consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
				"bootstrap.servers": cfg.BootstrapServers,
				"group.id":          cfg.ConsumerGroupID,
				"auto.offset.reset": "earliest",
			})
			if err != nil {
				return errors.Wrap(err, "failed to create Kafka consumer")
			}
			defer consumer.Close()

			err = consumer.Subscribe(cfg.Topic, nil)
			if err != nil {
				return errors.Wrap(err, "failed to subscribe to topic")
			}

			slog.Info("Consumer started, waiting for messages...")

			go func() {
				for {
					msg, err := consumer.ReadMessage(-1)
					if err != nil {
						slog.Error(errors.Wrap(err, "failed to read message").Error())
						continue
					}

					var u user.User
					if err := avro.Unmarshal(schema, msg.Value, &u); err != nil {
						continue
					}

					slog.Info("Received user:", slog.Any("user", u))
				}
			}()

			interrupt := cmdutils.InterruptChan()
			<-interrupt

			slog.Info("Shutting down the consumer...")

			return nil
		},
	}
}
