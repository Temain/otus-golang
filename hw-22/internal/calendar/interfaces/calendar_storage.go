package calendar

import (
	"time"

	e "github.com/Temain/otus-golang/hw-22/internal/calendar/entities"
)

type ICalendarStorage interface {
	List() ([]e.Event, error)
	Get(id int64) (*e.Event, error)
	Search(date time.Time) (*e.Event, error)
	Add(*e.Event) error
	Update(*e.Event) error
	Delete(int64) error
}
