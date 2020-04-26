package api

import (
	"net/http"
	"time"

	"github.com/Temain/otus-golang/hw-26/internal/configer"
	"github.com/Temain/otus-golang/hw-26/internal/domain"
	"github.com/Temain/otus-golang/hw-26/internal/domain/entities"
	"github.com/Temain/otus-golang/hw-26/internal/logger"
)

func StartHttpServer() error {
	cfg := configer.ReadConfig()
	log := logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	calendar := domain.NewMemoryCalendar()

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
		_ = calendar.Add(r.Context(), event)
		log.Println("added new event")
	})
	return http.ListenAndServe(cfg.HttpListen, nil)
}
