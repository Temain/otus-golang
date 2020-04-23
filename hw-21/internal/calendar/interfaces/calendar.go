package calendar

import (
	"time"

	c "github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
)

type Calendar interface {
	List() ([]c.Event, error)
	Search(date time.Time) (*c.Event, error)
	Add(*c.Event) error
	Update(*c.Event) error
	Delete(int64) error
}
