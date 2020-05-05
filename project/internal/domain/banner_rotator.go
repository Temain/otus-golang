package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/Temain/otus-golang/project/internal/domain/entities"

	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
	"github.com/Temain/otus-golang/project/internal/domain/storages"
)

type BannerRotator struct {
	rotationStorage interfaces.RotationStorage
}

func NewBannerRotator(db *sqlx.DB) (interfaces.Rotator, error) {
	storage, err := storages.NewPgRotationStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create rotation storage: %v", err)
	}
	return &BannerRotator{rotationStorage: storage}, nil
}

func (r *BannerRotator) Add(ctx context.Context, bannerId int64, slotId int64) error {
	found, err := r.rotationStorage.Get(ctx, bannerId, slotId)
	if err != nil {
		return err
	}
	if found != nil {
		return fmt.Errorf("banner %d already exists in rotation for slot %d", bannerId, slotId)
	}

	item := &entities.Rotation{
		BannerId:  bannerId,
		SlotId:    slotId,
		StartedAt: time.Now(),
	}
	err = r.rotationStorage.Add(ctx, item)
	if err != nil {
		return fmt.Errorf("error on add banner %d in rotation for slot %d: %v", bannerId, slotId, err)
	}

	return nil
}

func (r *BannerRotator) Delete(ctx context.Context, bannerId int64, slotId int64) error {
	event, err := r.rotationStorage.Get(ctx, bannerId, slotId)
	if err != nil {
		return err
	}
	if event == nil {
		return fmt.Errorf("banner %d not found in rotation for slot %d: %v", bannerId, slotId, err)
	}

	err = r.rotationStorage.Delete(ctx, bannerId, slotId)
	if err != nil {
		return fmt.Errorf("error on delete banner %d from rotation for slot %d: %v", bannerId, slotId, err)
	}

	return nil
}

func (r *BannerRotator) Click(ctx context.Context, bannerId int64, slotId int64, groupId int64) error {
	panic("implement me")
}

func (r *BannerRotator) Buyout(ctx context.Context, bannerId int64, slotId int64, groupId int64) error {
	panic("implement me")
}

func (r *BannerRotator) Get(ctx context.Context, slotId int64, groupId int64) (int64, error) {
	panic("implement me")
}
