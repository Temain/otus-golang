package main

import (
	"net/http"
	"time"

	c "github.com/temain/otus-golang/hw-15/internal/calendar"
	"github.com/temain/otus-golang/hw-15/pkg/configer"
	"github.com/temain/otus-golang/hw-15/pkg/logger"
)

func main() {
	cfg := configer.ReadConfig()
	log := logger.NewLogger(cfg.LogFile, cfg.LogLevel)
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
