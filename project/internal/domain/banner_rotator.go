package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/Temain/otus-golang/project/internal/domain/enums"

	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/algorithm"
	algentities "github.com/Temain/otus-golang/project/algorithm/entities"
	alginterfaces "github.com/Temain/otus-golang/project/algorithm/interfaces"
	"github.com/Temain/otus-golang/project/internal/domain/entities"
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
	"github.com/Temain/otus-golang/project/internal/domain/storages"
)

type BannerRotator struct {
	algorithm        alginterfaces.RotationAlgorithm
	bannerStorage    interfaces.BannerStorage
	slotStorage      interfaces.SlotStorage
	groupStorage     interfaces.GroupStorage
	rotationStorage  interfaces.RotationStorage
	statisticStorage interfaces.StatisticsStorage
}

func NewBannerRotator(db *sqlx.DB) (interfaces.Rotator, error) {
	bannerStorage, err := storages.NewPgBannerStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create banner storage: %v", err)
	}

	slotStorage, err := storages.NewPgSlotStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create slot storage: %v", err)
	}

	groupStorage, err := storages.NewPgGroupStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create group storage: %v", err)
	}

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
		bannerStorage:    bannerStorage,
		slotStorage:      slotStorage,
		groupStorage:     groupStorage,
		rotationStorage:  rotationStorage,
		statisticStorage: statisticStorage,
		algorithm:        algorithm,
	}
	return rotator, nil
}

func (r *BannerRotator) Add(ctx context.Context, bannerId int64, slotId int64) error {
	banner, err := r.bannerStorage.Get(ctx, bannerId)
	if err != nil {
		return fmt.Errorf("error on get banner %d from database: %v", bannerId, err)
	}
	if banner == nil {
		return fmt.Errorf("banner %d not found in database", bannerId)
	}

	slot, err := r.slotStorage.Get(ctx, slotId)
	if err != nil {
		return fmt.Errorf("error on get slot %d from database: %v", slotId, err)
	}
	if slot == nil {
		return fmt.Errorf("slot %d not found in database", slotId)
	}

	found, err := r.rotationStorage.Get(ctx, bannerId, slotId)
	if err != nil {
		return fmt.Errorf("error on get rotation item for bannner %d and slot %d: %v", bannerId, slotId, err)
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
		return fmt.Errorf("error on get rotation for banner %d and slot %d: %v", bannerId, slotId, err)
	}
	if event == nil {
		return fmt.Errorf("banner %d not found in rotation for slot %d", bannerId, slotId)
	}

	err = r.rotationStorage.Delete(ctx, bannerId, slotId)
	if err != nil {
		return fmt.Errorf("error on delete banner %d from rotation for slot %d: %v", bannerId, slotId, err)
	}

	return nil
}

func (r *BannerRotator) Get(ctx context.Context, slotId int64, groupId int64) (int64, error) {
	statistics, err := r.statisticStorage.SummaryBySlotAndGroup(ctx, slotId, groupId)
	if err != nil {
		return 0, err
	}

	var data []algentities.AlgorithmData
	for _, stat := range statistics {
		item := algentities.AlgorithmData{
			HandleId:  stat.BannerId,
			Count:     stat.Buyouts,
			AvgIncome: stat.Clicks,
		}
		data = append(data, item)
	}

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
	err := r.addStatistic(ctx, enums.Buyout, bannerId, slotId, groupId)
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
