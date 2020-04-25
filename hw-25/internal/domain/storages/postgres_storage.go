package storages

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/hw-25/internal/domain/entities"
	interfaces "github.com/Temain/otus-golang/hw-25/internal/domain/interfaces"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage(db *sqlx.DB) (interfaces.EventStorage, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db}, nil
}

func (pc *PostgresStorage) List(ctx context.Context) ([]entities.Event, error) {
	query := `
		SELECT id, title, description, created FROM events
	`

	var events []entities.Event
	err := pc.db.SelectContext(ctx, &events, query)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (pc *PostgresStorage) Search(ctx context.Context, created time.Time) (*entities.Event, error) {
	query := `
		SELECT id, title, description, created 
		FROM events 
		WHERE created = $1 
	`

	var event entities.Event
	err := pc.db.GetContext(ctx, &event, query, created)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (pc *PostgresStorage) Get(ctx context.Context, id int64) (*entities.Event, error) {
	query := `
		SELECT id, title, description, created 
		FROM events 
		WHERE id = $1 
	`

	var event entities.Event
	err := pc.db.GetContext(ctx, &event, query, id)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (pc *PostgresStorage) Add(ctx context.Context, event *entities.Event) error {
	query := `
		INSERT INTO events(title, description, created)
		VALUES (:title, :description, :created)
	`
	_, err := pc.db.NamedExecContext(ctx, query, event)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PostgresStorage) Update(ctx context.Context, event *entities.Event) error {
	query := `
		UPDATE events
		SET title = :title,
			description = :description,
			created = :created
		where id = :id
	`
	_, err := pc.db.NamedExecContext(ctx, query, event)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PostgresStorage) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE 
		FROM events
		where id = $1
	`
	_, err := pc.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
