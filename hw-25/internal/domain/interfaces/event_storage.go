package calendar

import (
	"context"
	"time"

	"github.com/Temain/otus-golang/hw-25/internal/domain/entities"
)

type EventStorage interface {
	List(ctx context.Context) ([]entities.Event, error)
	Get(ctx context.Context, id int64) (*entities.Event, error)
	Search(ctx context.Context, date time.Time) (*entities.Event, error)
	Add(ctx context.Context, event *entities.Event) error
	Update(ctx context.Context, event *entities.Event) error
	Delete(ctx context.Context, id int64) error
}
