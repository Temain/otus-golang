package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"github.com/spf13/pflag"
	"github.com/temain/otus-golang/hw-15/configs"
	c "github.com/temain/otus-golang/hw-15/pkg/calendar"
)

func main() {
	cfg := readConfig()
	calendar := c.NewCalendar()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello")
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		event := &c.Event{
			Id:          1,
			Title:       "Morning coffee",
			Description: "The most important event of the day",
			Date:        time.Now(),
		}
		_ = calendar.Add(event)
		log.Println("added new event")
	})
	err := http.ListenAndServe(cfg.HttpListen, nil)
	if err != nil {
		log.Fatalf("http server error: %v", err)
	}
}

func readConfig() *configs.Config {
	var configPath string
	pflag.StringVarP(&configPath, "config", "c", "configs/config.json", "Config file path")

	loader := confita.NewLoader(
		file.NewBackend(configPath),
		flags.NewBackend(),
	)
	cfg := configs.Config{}
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	return &cfg
}

func configureLogging() {

}
