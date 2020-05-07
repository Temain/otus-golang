package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/Temain/otus-golang/project/internal/domain/enums"

	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
	"github.com/Temain/otus-golang/project/internal/domain/storages"
	"github.com/Temain/otus-golang/project/pkg/algorithm"
	alg_entities "github.com/Temain/otus-golang/project/pkg/algorithm/entities"
	alg_interfaces "github.com/Temain/otus-golang/project/pkg/algorithm/interfaces"
)

type BannerRotator struct {
	algorithm        alg_interfaces.RotationAlgorithm
	rotationStorage  interfaces.RotationStorage
	statisticStorage interfaces.StatisticsStorage
}

func NewBannerRotator(db *sqlx.DB) (interfaces.Rotator, error) {
	rotationStorage, err := storages.NewPgRotationStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create rotation storage: %v", err)
	}

	statisticStorage, err := storages.NewPgStatisticsStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create statistic storage: %v", err)
	}

	algorithm, err := algorithm.NewMultiarmedBandit()
	if err != nil {
		return nil, fmt.Errorf("error on init rotation algorithm")
	}

	rotator := &BannerRotator{
		rotationStorage:  rotationStorage,
		statisticStorage: statisticStorage,
		algorithm:        algorithm,
	}
	return rotator, nil
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

func (r *BannerRotator) Get(ctx context.Context, slotId int64, groupId int64) (int64, error) {
	_, err := r.statisticStorage.SummaryBySlotAndGroup(ctx, slotId, groupId)
	if err != nil {
		return 0, err
	}

	var data []alg_entities.AlgorithmData
	return r.algorithm.GetHandle(data)
}

func (r *BannerRotator) Click(ctx context.Context, bannerId int64, slotId int64, groupId int64) error {
	err := r.addStatistic(ctx, enums.Click, bannerId, slotId, groupId)
	if err != nil {
		return fmt.Errorf("error on save statistic by click: %v", err)
	}

	return nil
}

func (r *BannerRotator) Buyout(ctx context.Context, bannerId int64, slotId int64, groupId int64) error {
	err := r.addStatistic(ctx, enums.Click, bannerId, slotId, groupId)
	if err != nil {
		return fmt.Errorf("error on save statistic by buyout: %v", err)
	}

	return nil
}

func (r *BannerRotator) addStatistic(ctx context.Context, statType enums.StatisticType, bannerId int64, slotId int64, groupId int64) error {
	item := &entities.Statistics{
		TypeId:   int64(statType),
		BannerId: bannerId,
		SlotId:   slotId,
		GroupId:  groupId,
		DateTime: time.Now(),
	}

	err := r.statisticStorage.Add(ctx, item)
	if err != nil {
		return fmt.Errorf("error on save statistic, banner %d slot %d group %d: %v", bannerId, slotId, groupId, err)
	}

	return nil
}
