package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

var SchedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "run scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running scheduler...")

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

		body := "Hello World!"
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "failed to publish a message")
	},
}

func getEvents() {

}
