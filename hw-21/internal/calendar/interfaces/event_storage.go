package calendar

import (
	"time"

	"github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
)

type EventStorage interface {
	List() ([]entities.Event, error)
	Get(id int64) (*entities.Event, error)
	Search(date time.Time) (*entities.Event, error)
	Add(*entities.Event) error
	Update(*entities.Event) error
	Delete(int64) error
}
