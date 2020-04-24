package cmd

import (
	"fmt"
	"net/http"
	"time"

	c "github.com/Temain/otus-golang/hw-22/internal/calendar"
	e "github.com/Temain/otus-golang/hw-22/internal/calendar/entities"
	"github.com/Temain/otus-golang/hw-22/internal/configer"
	"github.com/Temain/otus-golang/hw-22/internal/logger"

	"github.com/spf13/cobra"
)

var HttpServerCmd = &cobra.Command{
	Use:   "http_server",
	Short: "run http server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running http server...")

		cfg := configer.ReadConfig()
		log := logger.NewLogger(cfg.LogFile, cfg.LogLevel)
		calendar := c.NewMemoryCalendar()

		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			log.Println("hello")
		})

		http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
			event := &e.Event{
				Id:          1,
				Title:       "Morning coffee",
				Description: "The most important event of the day",
				Created:     time.Now(),
			}
			_ = calendar.Add(event)
			log.Println("added new event")
		})
		err := http.ListenAndServe(cfg.HttpListen, nil)
		if err != nil {
			log.Fatalf("http server error: %v", err)
		}
	},
}
