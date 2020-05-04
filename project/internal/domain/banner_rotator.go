package domain

import (
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
	"github.com/jmoiron/sqlx"
)

type BannerRotator struct {
}

func NewBannerRotator(db *sqlx.DB) (interfaces.Rotator, error) {
	/*storage, err := storages.NewPostgresStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create postgres storage: %v", err)
	}*/
	return &BannerRotator{}, nil
}
