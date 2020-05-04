package interfaces

import (
	"context"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
)

type BannerStorage interface {
	List(ctx context.Context) ([]entities.Banner, error)
	Get(ctx context.Context, id int64) (*entities.Banner, error)
	Search(ctx context.Context, pattern string) (*entities.Banner, error)
	Add(ctx context.Context, event *entities.Banner) error
	Update(ctx context.Context, event *entities.Banner) error
	Delete(ctx context.Context, id int64) error
}
