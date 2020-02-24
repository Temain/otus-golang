package main

import (
	"log"
	"net/http"
	"time"

	c "github.com/temain/otus-golang/hw-15/pkg/calendar"
)

func main() {
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
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("http server error: %v", err)
	}
}
