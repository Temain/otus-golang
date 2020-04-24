package calendar

import (
	"context"
	"time"

	c "github.com/Temain/otus-golang/hw-22/internal/calendar/entities"
)

type ICalendar interface {
	List(ctx context.Context) ([]c.Event, error)
	Search(ctx context.Context, date time.Time) (*c.Event, error)
	Add(ctx context.Context, event *c.Event) error
	Update(ctx context.Context, event *c.Event) error
	Delete(ctx context.Context, id int64) error
}
