package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Temain/otus-golang/hw-25/internal/rabbitmq"

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
		ctx, _ := withSignal(context.Background(), os.Interrupt)
		consumer := rabbitmq.NewConsumer(ctx, "tag", cfg.RabbitUrl, cfg.RabbitExchange, cfg.RabbitExchangeType, cfg.RabbitQueue, "")
		err := consumer.Handle(handleMessages, 5)
		if err != nil {
			log.Fatalf("error on handle messages: %v", err)
		}

		select {
		case <-ctx.Done():
			fmt.Println("shutdown signal received")
		}
	},
}

func handleMessages(ctx context.Context, delivery <-chan amqp.Delivery) {
	for {
		select {
		case d := <-delivery:
			log.Printf("received a message: %s", d.Body)
			break
		case <-ctx.Done():
			log.Printf("goroutine done")
			return
		}
	}
}

func withSignal(ctx context.Context, s ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, s...)
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
			cancel()
		}
		signal.Stop(c)
	}()
	return ctx, cancel
}
