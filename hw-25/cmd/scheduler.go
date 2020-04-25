package cmd

import (
	"context"
	"encoding/json"
	"log"
	"time"

	i "github.com/Temain/otus-golang/hw-25/internal/domain/interfaces"

	e "github.com/Temain/otus-golang/hw-25/internal/domain/entities"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"

	"github.com/Temain/otus-golang/hw-25/internal/configer"
	s "github.com/Temain/otus-golang/hw-25/internal/domain/storages"
)

var SchedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "run scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running scheduler...")

		cfg := configer.ReadConfig()
		conn, err := amqp.Dial(cfg.RabbitUrl)
		failOnError(err, "failed to connect to RabbitMQ")
		defer conn.Close()

		channel, err := conn.Channel()
		failOnError(err, "failed to open a channel")
		defer channel.Close()

		queue, err := channel.QueueDeclare(
			cfg.RabbitQueue,
			false,
			false,
			false,
			false,
			nil,
		)
		failOnError(err, "failed to declare a queue")

		ctx := context.Background()
		storage, err := s.NewPostgreStorage(cfg.PostgreDSN)
		if err != nil {
			failOnError(err, "connection to database failed")
		}

		duration := 10 * time.Second
		log.Printf("check events every %v", duration)
		for range time.Tick(duration) {
			sendMessage(ctx, storage, channel, queue)
		}
	},
}

func sendMessage(ctx context.Context, storage i.ICalendarStorage, ch *amqp.Channel, q amqp.Queue) {
	events, err := getEvents(ctx, storage)
	if err != nil {
		failOnError(err, "error on get events")
	}

	for _, event := range events {
		body, err := json.Marshal(event)
		if err != nil {
			failOnError(err, "error on marshal event")
		}
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf("sent %s\n", body)
		failOnError(err, "failed to publish a message")
	}
}

func getEvents(ctx context.Context, storage i.ICalendarStorage) ([]e.Event, error) {
	// TODO: additional logic for select events
	events, err := storage.List(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil
}
