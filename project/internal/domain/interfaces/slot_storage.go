package interfaces

import (
	"context"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
)

type SlotStorage interface {
	List(ctx context.Context) ([]entities.Slot, error)
	Get(ctx context.Context, id int64) (*entities.Slot, error)
	Search(ctx context.Context, pattern string) (*entities.Slot, error)
	Add(ctx context.Context, event *entities.Slot) error
	Update(ctx context.Context, event *entities.Slot) error
	Delete(ctx context.Context, id int64) error
}
