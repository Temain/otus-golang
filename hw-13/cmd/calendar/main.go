package main

import (
	"time"

	c "github.com/temain/otus-golang/hw-13/pkg/calendar"
)

func main() {
	var calendar c.Calendar
	calendar = c.NewCalendar()

	// Добавление нового события.
	event := &c.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Date:        time.Now(),
	}
	_ = calendar.Add(event)
}
