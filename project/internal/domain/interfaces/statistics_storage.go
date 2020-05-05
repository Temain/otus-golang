package interfaces

import (
	"context"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
)

type StatisticsStorage interface {
	// List(ctx context.Context, typeId int64, bannerId int64, slotId int64, groupId int64) ([]entities.Statistics, error)
	Add(ctx context.Context, item *entities.Statistics) error
}
