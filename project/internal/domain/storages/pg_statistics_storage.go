package storages

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
)

type PgStatisticsStorage struct {
	db *sqlx.DB
}

func NewPgStatisticsStorage(db *sqlx.DB) (interfaces.StatisticsStorage, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &PgStatisticsStorage{db}, nil
}

func (pc *PgStatisticsStorage) Add(ctx context.Context, item *entities.Statistics) error {
	query := `
		INSERT INTO statistics(type_id, banner_id, slot_id, group_id, date_time)
		VALUES (:type_id, :banner_id, :slot_id, :group_id, :date_time)
	`
	_, err := pc.db.NamedExecContext(ctx, query, item)
	if err != nil {
		return err
	}

	return nil
}
