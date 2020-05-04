package domain

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
	"github.com/Temain/otus-golang/project/internal/domain/storages"
)

type BannerRotator struct {
	storage interfaces.RotationStorage
}

func NewBannerRotator(db *sqlx.DB) (interfaces.Rotator, error) {
	storage, err := storages.NewPgRotationStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create rotation storage: %v", err)
	}
	return &BannerRotator{storage: storage}, nil
}
