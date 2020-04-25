package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v3"
	"github.com/streadway/amqp"
)

type Consumer struct {
	ctx          context.Context
	conn         *amqp.Connection
	channel      *amqp.Channel
	done         chan error
	consumerTag  string
	uri          string
	exchangeName string
	exchangeType string
	queue        string
	bindingKey   string
}

func NewConsumer(ctx context.Context, consumerTag, uri, exchangeName, exchangeType, queue, bindingKey string) *Consumer {
	return &Consumer{
		ctx:          ctx,
		consumerTag:  consumerTag,
		uri:          uri,
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		queue:        queue,
		bindingKey:   bindingKey,
		done:         make(chan error),
	}
}

func (c *Consumer) reConnect() (<-chan amqp.Delivery, error) {
	be := backoff.NewExponentialBackOff()
	be.MaxElapsedTime = time.Minute
	be.InitialInterval = 1 * time.Second
	be.Multiplier = 2
	be.MaxInterval = 15 * time.Second

	b := backoff.WithContext(be, context.Background())
	for {
		d := b.NextBackOff()
		if d == backoff.Stop {
			return nil, fmt.Errorf("stop reconnecting")
		}

		select {
		case <-time.After(d):
			if err := c.connect(); err != nil {
				log.Printf("could not connect in reconnect call: %+v", err)
				continue
			}
			msgs, err := c.announceQueue()
			if err != nil {
				fmt.Printf("Couldn't connect: %+v", err)
				continue
			}

			return msgs, nil
		}
	}
}

func (c *Consumer) connect() error {
	var err error
	c.conn, err = amqp.Dial(c.uri)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %s", err)
	}

	go func() {
		log.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
		// Понимаем, что канал сообщений закрыт, надо пересоздать соединение.
		c.done <- errors.New("channel Closed")
	}()

	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	if err = c.channel.ExchangeDeclare(
		c.exchangeName,
		c.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("exchange declare: %s", err)
	}

	return nil
}

// Задекларировать очередь, которую будем слушать.
func (c *Consumer) announceQueue() (<-chan amqp.Delivery, error) {
	queue, err := c.channel.QueueDeclare(
		c.queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("queue Declare: %s", err)
	}

	// Число сообщений, которые можно подтвердить за раз.
	err = c.channel.Qos(50, 0, false)
	if err != nil {
		return nil, fmt.Errorf("error setting qos: %s", err)
	}

	// Создаём биндинг (правило маршрутизации).
	if err = c.channel.QueueBind(
		queue.Name,
		c.bindingKey,
		c.exchangeName,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("queue Bind: %s", err)
	}

	msgs, err := c.channel.Consume(
		queue.Name,
		c.consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("queue Consume: %s", err)
	}

	return msgs, nil
}

func (c *Consumer) Handle(fn func(ctx context.Context, delivery <-chan amqp.Delivery), threads int) error {
	var err error
	if err = c.connect(); err != nil {
		return fmt.Errorf("error: %v", err)
	}
	delivery, err := c.announceQueue()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("waiting for messages...")

	wg := sync.WaitGroup{}

	for {
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go func(ctx context.Context, delivery <-chan amqp.Delivery) {
				fn(ctx, delivery)
				wg.Done()
			}(c.ctx, delivery)
		}

		select {
		case err := <-c.done:
			if err != nil {
				delivery, err = c.reConnect()
				if err != nil {
					return fmt.Errorf("reconnecting Error: %s", err)
				}
				fmt.Println("reconnected... possibly")
			}
			break
		case <-c.ctx.Done():
			wg.Wait()
			fmt.Println("handle done")
			return nil
		}
	}
}
