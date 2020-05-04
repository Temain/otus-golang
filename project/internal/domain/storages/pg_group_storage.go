package storages

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
)

type PgGroupStorage struct {
	db *sqlx.DB
}

func NewPgGroupStorage(db *sqlx.DB) (interfaces.GroupStorage, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &PgGroupStorage{db}, nil
}

func (pc *PgGroupStorage) List(ctx context.Context) ([]entities.Group, error) {
	query := `
		SELECT id, title FROM groups
	`

	var groups []entities.Group
	err := pc.db.SelectContext(ctx, &groups, query)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (pc *PgGroupStorage) Search(ctx context.Context, pattern string) (*entities.Group, error) {
	query := `
		SELECT id, title 
		FROM groups 
		WHERE title LIKE '%' || $1 || '%'
	`

	var group entities.Group
	err := pc.db.GetContext(ctx, &group, query, pattern)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &group, nil
}

func (pc *PgGroupStorage) Get(ctx context.Context, id int64) (*entities.Group, error) {
	query := `
		SELECT id, title 
		FROM groups 
		WHERE id = $1 
	`

	var group entities.Group
	err := pc.db.GetContext(ctx, &group, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &group, nil
}

func (pc *PgGroupStorage) Add(ctx context.Context, group *entities.Group) error {
	query := `
		INSERT INTO groups(title)
		VALUES (:title)
	`
	_, err := pc.db.NamedExecContext(ctx, query, group)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PgGroupStorage) Update(ctx context.Context, group *entities.Group) error {
	query := `
		UPDATE groups
		SET title = :title
		where id = :id
	`
	_, err := pc.db.NamedExecContext(ctx, query, group)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PgGroupStorage) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE 
		FROM groups
		where id = $1
	`
	_, err := pc.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
