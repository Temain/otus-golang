package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/streadway/amqp"
)

var SenderCmd = &cobra.Command{
	Use:   "sender",
	Short: "run sender",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running sender...")

		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"events", // name
			false,    // durable
			false,    // delete when unused
			false,    // exclusive
			false,    // no-wait
			nil,      // arguments
		)
		failOnError(err, "failed to declare a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
		}()

		log.Printf(" [*] waiting for messages. To exit press CTRL+C")
		<-forever
	},
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
