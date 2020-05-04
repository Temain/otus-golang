package interfaces

import (
	"context"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
)

type GroupStorage interface {
	List(ctx context.Context) ([]entities.Group, error)
	Get(ctx context.Context, id int64) (*entities.Group, error)
	Search(ctx context.Context, pattern string) (*entities.Group, error)
	Add(ctx context.Context, event *entities.Group) error
	Update(ctx context.Context, event *entities.Group) error
	Delete(ctx context.Context, id int64) error
}
