package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/pflag"

	"github.com/Temain/otus-golang/hw-28/internal/configer"
	"github.com/Temain/otus-golang/hw-28/internal/domain/entities"
	interfaces "github.com/Temain/otus-golang/hw-28/internal/domain/interfaces"
	"github.com/Temain/otus-golang/hw-28/internal/domain/storages"
	"github.com/Temain/otus-golang/hw-28/internal/rabbitmq"
	"github.com/jmoiron/sqlx"
)

var configPath string

func init() {
	pflag.StringVarP(&configPath, "config", "c", "configs/config.json", "Config file path")
	pflag.Parse()
}

func main() {
	log.Println("running scheduler...")

	cfg := configer.ReadConfigScheduler(configPath)
	db, err := sqlx.Open("pgx", cfg.PostgresDsn)
	if err != nil {
		log.Fatalf("connection to database failed %v", err)
	}
	storage, err := storages.NewPostgresStorage(db)
	if err != nil {
		log.Fatalf("unable to create postgres storage: %v", err)
	}

	ctx := context.Background()
	producer := rabbitmq.NewProducer(ctx, cfg.RabbitUrl, cfg.RabbitExchange, cfg.RabbitQueue)

	duration := 10 * time.Second
	log.Printf("check events every %v", duration)
	for range time.Tick(duration) {
		sendMessage(ctx, storage, producer)
	}
}

func sendMessage(ctx context.Context, storage interfaces.EventStorage, producer *rabbitmq.Producer) {
	events, err := getEvents(ctx, storage)
	if err != nil {
		log.Fatalf("error on get events: %v", err)
	}

	for _, event := range events {
		body, err := json.Marshal(event)
		if err != nil {
			log.Fatalf("error on marshal event: %v", err)
		}

		err = producer.Publish(body)
		log.Printf("sent %s\n", body)
		if err != nil {
			log.Fatalf("failed to publish a message: %v", err)
		}
	}
}

func getEvents(ctx context.Context, storage interfaces.EventStorage) ([]entities.Event, error) {
	events, err := storage.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error on retrieve events: %v", err)
	}

	return events, nil
}
