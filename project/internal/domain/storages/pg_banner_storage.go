package storages

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/project/internal/domain/entities"
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
)

type PgBannerStorage struct {
	db *sqlx.DB
}

func NewPgBannerStorage(db *sqlx.DB) (interfaces.BannerStorage, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &PgBannerStorage{db}, nil
}

func (pc *PgBannerStorage) List(ctx context.Context) ([]entities.Banner, error) {
	query := `
		SELECT id, title FROM banners
	`

	var banners []entities.Banner
	err := pc.db.SelectContext(ctx, &banners, query)
	if err != nil {
		return nil, err
	}

	return banners, nil
}

func (pc *PgBannerStorage) Search(ctx context.Context, pattern string) (*entities.Banner, error) {
	query := `
		SELECT id, title 
		FROM banners 
		WHERE title LIKE '%' || $1 || '%'
	`

	var banner entities.Banner
	err := pc.db.GetContext(ctx, &banner, query, pattern)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &banner, nil
}

func (pc *PgBannerStorage) Get(ctx context.Context, id int64) (*entities.Banner, error) {
	query := `
		SELECT id, title 
		FROM banners 
		WHERE id = $1 
	`

	var banner entities.Banner
	err := pc.db.GetContext(ctx, &banner, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &banner, nil
}

func (pc *PgBannerStorage) Add(ctx context.Context, banner *entities.Banner) error {
	query := `
		INSERT INTO banners(title)
		VALUES (:title)
	`
	_, err := pc.db.NamedExecContext(ctx, query, banner)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PgBannerStorage) Update(ctx context.Context, banner *entities.Banner) error {
	query := `
		UPDATE banners
		SET title = :title
		where id = :id
	`
	_, err := pc.db.NamedExecContext(ctx, query, banner)
	if err != nil {
		return err
	}

	return nil
}

func (pc *PgBannerStorage) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE 
		FROM banners
		where id = $1
	`
	_, err := pc.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
