package interfaces

import (
	"context"
)

type Rotator interface {
	Add(ctx context.Context, bannerId int64, slotId int64) error
	Delete(ctx context.Context, bannerId int64, slotId int64) error
	Click(ctx context.Context, bannerId int64, slotId int64, groupId int64) error
	Buyout(ctx context.Context, bannerId int64, slotId int64, groupId int64) error
	Get(ctx context.Context, slotId int64, groupId int64) (int64, error)
}
