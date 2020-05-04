package storages

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
)

type PgSlotStorage struct {
	db *sqlx.DB
}

func NewPgSlotStorage(db *sqlx.DB) (interfaces.SlotStorage, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &PgSlotStorage{db}, nil
}

func (pc *PgSlotStorage) List(ctx context.Context) ([]entities.Slot, error) {
	query := `
		SELECT id, title FROM slots
	`

	var slots []entities.Slot
	err := pc.db.SelectContext(ctx, &slots, query)
	if err != nil {
		return nil, err
	}

	return slots, nil
}

func (pc *PgSlotStorage) Search(ctx context.Context, pattern string) (*entities.Slot, error) {
	query := `
		SELECT id, title 
		FROM slots 
		WHERE title LIKE '%' || $1 || '%'
	`

	var slot entities.Slot
	err := pc.db.GetContext(ctx, &slot, query, pattern)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &slot, nil
}

func (pc *PgSlotStorage) Get(ctx context.Context, id int64) (*entities.Slot, error) {
	query := `
		SELECT id, title 
		FROM slots 
		WHERE id = $1 
	`

	var slot entities.Slot
	err := pc.db.GetContext(ctx, &slot, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &slot, nil
}

func (pc *PgSlotStorage) Add(ctx context.Context, slot *entities.Slot) error {
	query := `
		INSERT INTO slots(title)
		VALUES (:title)
	`
	_, err := pc.db.NamedExecContext(ctx, query, slot)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PgSlotStorage) Update(ctx context.Context, slot *entities.Slot) error {
	query := `
		UPDATE slots
		SET title = :title
		where id = :id
	`
	_, err := pc.db.NamedExecContext(ctx, query, slot)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PgSlotStorage) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE 
		FROM slots
		where id = $1
	`
	_, err := pc.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
