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

func (pc *PgStatisticsStorage) SummaryBySlotAndGroup(ctx context.Context, slotId int64, groupId int64) ([]entities.StatisticsSummary, error) {
	query := `
		SELECT r.banner_id, r.slot_id, sb.group_id,
			COALESCE(sb.buyouts, 0) as buyouts, 
			COALESCE(sc.clicks, 0) AS clicks
		FROM rotation r 
		LEFT JOIN (
			select banner_id, group_id, count(*) as buyouts
			from "statistics" s 
			where slot_id = $1 and group_id = $2 AND type_id = 1
			group by type_id, banner_id, slot_id, group_id
		) AS sb ON sb.banner_id = r.banner_id
		LEFT JOIN (
			select banner_id, count(*) as clicks
			from "statistics" s 
			where slot_id = $1 and group_id = $2 AND type_id = 2
			group by type_id, banner_id, slot_id, group_id
		) AS sc ON sc.banner_id = r.banner_id
		WHERE slot_id = $1
	`

	var events []entities.StatisticsSummary
	err := pc.db.SelectContext(ctx, &events, query, slotId, groupId)
	if err != nil {
		return nil, err
	}

	return events, nil
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
