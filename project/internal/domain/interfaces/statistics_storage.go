package interfaces

import (
	"context"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
)

type StatisticsStorage interface {
	SummaryBySlotAndGroup(ctx context.Context, slotId int64, groupId int64) ([]entities.StatisticsSummary, error)
	Add(ctx context.Context, item *entities.Statistics) error
}
