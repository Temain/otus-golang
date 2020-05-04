package interfaces

import (
	"context"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
)

type RotationStorage interface {
	List(ctx context.Context) ([]entities.Rotation, error)
	ListBySlot(ctx context.Context, slotId int64) ([]entities.Rotation, error)
	Get(ctx context.Context, bannerId int64) (*entities.Rotation, error)
	Add(ctx context.Context, event *entities.Rotation) error
	Update(ctx context.Context, event *entities.Rotation) error
	Delete(ctx context.Context, id int64) error
}
