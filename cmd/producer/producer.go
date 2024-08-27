package producer

import (
	"log/slog"
	"time"

	"avro/pkg/cmdutils"
	"avro/pkg/config"
	"avro/pkg/gen/go/user"
	"avro/pkg/types"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"github.com/hamba/avro"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cfg = config.Producer{
	SchemaPaths:      []string{"/usr/local/lib/schema/avro/user.avsc"},
	BootstrapServers: "kafka:29092",
	Topic:            "kita.user",
}

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "producer",
		RunE: func(cmd *cobra.Command, args []string) error {
			schema, err := avro.ParseFiles(cfg.SchemaPaths...)
			if err != nil {
				return errors.Wrap(err, "failed to parse Avro schema")
			}

			producer, err := kafka.NewProducer(&kafka.ConfigMap{
				"bootstrap.servers": cfg.BootstrapServers,
			})
			if err != nil {
				return errors.Wrap(err, "failed to create Kafka producer")
			}
			defer producer.Close()

			slog.Info("Producer started, generating messages...")

			go func(pchan <-chan kafka.Event) {
				for {
					u := generateRandomUser()

					slog.Info("Generated user:", slog.Any("user", u))

					userData, err := avro.Marshal(schema, u)
					if err != nil {
						slog.Error(errors.Wrap(err, "failed to marshal user data").Error())
						return
					}

					if err := producer.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{
							Topic:     types.ToPtr(cfg.Topic),
							Partition: kafka.PartitionAny,
						},
						Value: userData,
					}, nil); err != nil {
						slog.Error(errors.Wrap(err, "failed to produce message").Error())
						return
					}

					select {
					case e, ok := <-pchan:
						if !ok {
							return
						}
						if msg, ok := e.(*kafka.Message); ok {
							if msg.TopicPartition.Error != nil {
								slog.Error("Delivery failed", slog.String("partition", msg.TopicPartition.String()))
							} else {
								slog.Info("Delivered message", slog.String("partition", msg.TopicPartition.String()))
							}
						}
					case <-time.After(5 * time.Second):
						return
					}
				}
			}(producer.Events())

			interrupt := cmdutils.InterruptChan()
			<-interrupt

			slog.Info("Shutting down the application...")

			return nil
		},
	}
}

func generateRandomUser() user.User {
	return user.User{
		ID:    uuid.NewString(),
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Age:   types.ToPtr(10),
	}
}
