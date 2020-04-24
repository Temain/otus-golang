package storages

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"

	e "github.com/Temain/otus-golang/hw-22/internal/calendar/entities"
	i "github.com/Temain/otus-golang/hw-22/internal/calendar/interfaces"
)

type PostgreStorage struct {
	db *sqlx.DB
}

func NewPostgreStorage(dsn string) (i.ICalendarStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgreStorage{db}, nil
}
func CreatePostgreStorage(db *sqlx.DB) (i.ICalendarStorage, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgreStorage{db}, nil
}

func (pc *PostgreStorage) List(ctx context.Context) ([]e.Event, error) {
	query := `
		SELECT id, title, description, created FROM events
	`

	var events []e.Event
	err := pc.db.SelectContext(ctx, &events, query)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (pc *PostgreStorage) Search(ctx context.Context, created time.Time) (*e.Event, error) {
	query := `
		SELECT id, title, description, created 
		FROM events 
		WHERE created = $1 
	`

	var event e.Event
	err := pc.db.GetContext(ctx, &event, query, created)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (pc *PostgreStorage) Get(ctx context.Context, id int64) (*e.Event, error) {
	query := `
		SELECT id, title, description, created 
		FROM events 
		WHERE id = $1 
	`

	var event e.Event
	err := pc.db.GetContext(ctx, &event, query, id)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (pc *PostgreStorage) Add(ctx context.Context, event *e.Event) error {
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

func (pc *PostgreStorage) Update(ctx context.Context, event *e.Event) error {
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

func (pc *PostgreStorage) Delete(ctx context.Context, id int64) error {
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
