package api

import (
	"net/http"
	"time"

	"github.com/Temain/otus-golang/hw-21/internal/calendar"
	"github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
	"github.com/Temain/otus-golang/hw-21/internal/configer"
	"github.com/Temain/otus-golang/hw-21/internal/logger"
)

func StartHttpServer() error {
	cfg := configer.ReadConfig()
	log := logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	calendar := calendar.NewMemoryCalendar()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello")
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		event := &entities.Event{
			Id:          1,
			Title:       "Morning coffee",
			Description: "The most important event of the day",
			Created:     time.Now(),
		}
		_ = calendar.Add(event)
		log.Println("added new event")
	})
	return http.ListenAndServe(cfg.HttpListen, nil)
}
