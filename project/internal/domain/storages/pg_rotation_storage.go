package storages

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
)

type PgRotationStorage struct {
	db *sqlx.DB
}

func NewPgRotationStorage(db *sqlx.DB) (interfaces.RotationStorage, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &PgRotationStorage{db}, nil
}

func (pc *PgRotationStorage) List(ctx context.Context) ([]entities.Rotation, error) {
	query := `
		SELECT id, title FROM rotation
	`

	var list []entities.Rotation
	err := pc.db.SelectContext(ctx, &list, query)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (pc *PgRotationStorage) ListBySlot(ctx context.Context, slotId int64) ([]entities.Rotation, error) {
	query := `
		SELECT banner_id, slot_id, started_at 
		FROM rotation
		WHERE slot_id = $1
	`

	var list []entities.Rotation
	err := pc.db.SelectContext(ctx, &list, query, slotId)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (pc *PgRotationStorage) Get(ctx context.Context, id int64) (*entities.Rotation, error) {
	query := `
		SELECT banner_id, slot_id, started_at 
		FROM rotation 
		WHERE id = $1 
	`

	var item entities.Rotation
	err := pc.db.GetContext(ctx, &item, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (pc *PgRotationStorage) Add(ctx context.Context, item *entities.Rotation) error {
	query := `
		INSERT INTO rotation(banner_id, slot_id, started_at)
		VALUES (:bannerId, :slotId, :startedAt)
	`
	_, err := pc.db.NamedExecContext(ctx, query, item)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PgRotationStorage) Update(ctx context.Context, item *entities.Rotation) error {
	query := `
		UPDATE rotation
		SET banner_id = :banner_id,
		    slot_id = :slot_id,
		    started_at = :started_at
		where id = :id
	`
	_, err := pc.db.NamedExecContext(ctx, query, item)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PgRotationStorage) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE 
		FROM rotation
		where id = $1
	`
	_, err := pc.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
