package rabbitmq

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"
)

type Producer struct {
	ctx          context.Context
	conn         *amqp.Connection
	channel      *amqp.Channel
	uri          string
	exchangeName string
	queue        string
}

func NewProducer(ctx context.Context, uri string, exchangeName, queue string) *Producer {
	return &Producer{
		ctx:          ctx,
		uri:          uri,
		exchangeName: exchangeName,
		queue:        queue,
	}
}

func (p *Producer) connect() error {
	var err error
	p.conn, err = amqp.Dial(p.uri)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %s", err)
	}

	p.channel, err = p.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}

	return nil
}

func (p *Producer) Publish(body []byte) error {
	var err error
	if err = p.connect(); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	err = p.channel.Publish(
		p.exchangeName,
		p.queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err
}
