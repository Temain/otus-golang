package cmd

import (
	"log"

	"github.com/Temain/otus-golang/hw-25/internal/configer"
	"github.com/spf13/cobra"

	"github.com/streadway/amqp"
)

var SenderCmd = &cobra.Command{
	Use:   "sender",
	Short: "run sender",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running sender...")

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

		msgs, err := channel.Consume(
			queue.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		failOnError(err, "failed to register a consumer")

		forever := make(chan bool)
		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
		}()

		log.Printf("waiting for messages...")
		<-forever
	},
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
