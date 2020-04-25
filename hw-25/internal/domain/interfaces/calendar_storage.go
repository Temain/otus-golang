package calendar

import (
	"context"
	"time"

	e "github.com/Temain/otus-golang/hw-25/internal/domain/entities"
)

type ICalendarStorage interface {
	List(ctx context.Context) ([]e.Event, error)
	Get(ctx context.Context, id int64) (*e.Event, error)
	Search(ctx context.Context, date time.Time) (*e.Event, error)
	Add(ctx context.Context, event *e.Event) error
	Update(ctx context.Context, event *e.Event) error
	Delete(ctx context.Context, id int64) error
}
