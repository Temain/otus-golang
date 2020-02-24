package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/temain/otus-golang/hw-15/configs"
	c "github.com/temain/otus-golang/hw-15/pkg/calendar"
)

func main() {
	cfg := readConfig()
	configLogging(cfg)

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

func configLogging(cfg *configs.Config) {
	f, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	var level log.Level
	switch cfg.LogLevel {
	case "info":
		level = log.InfoLevel
		break
	case "debug":
		level = log.DebugLevel
		break
	case "warn":
		level = log.WarnLevel
		break
	case "error":
		level = log.ErrorLevel
		break
	}
	log.SetLevel(level)
}
