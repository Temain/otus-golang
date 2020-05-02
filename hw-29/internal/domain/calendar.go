package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/Temain/otus-golang/hw-29/internal/domain/entities"
	interfaces "github.com/Temain/otus-golang/hw-29/internal/domain/interfaces"
	"github.com/Temain/otus-golang/hw-29/internal/domain/storages"
)

type Calendar struct {
	storage interfaces.EventStorage
}

func NewMemoryCalendar() interfaces.EventAdapter {
	storage := storages.NewMemoryStorage()
	return &Calendar{storage}
}
func NewPostgresCalendar(db *sqlx.DB) (interfaces.EventAdapter, error) {
	storage, err := storages.NewPostgresStorage(db)
	if err != nil {
		return nil, fmt.Errorf("unable to create postgres storage: %v", err)
	}
	return &Calendar{storage}, nil
}
func CreateCalendar(storage interfaces.EventStorage) interfaces.EventAdapter {
	return &Calendar{storage}
}

func (c *Calendar) List(ctx context.Context) ([]entities.Event, error) {
	list, err := c.storage.List(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Calendar) Search(ctx context.Context, created time.Time) (*entities.Event, error) {
	found, err := c.storage.Search(ctx, created)
	if err != nil {
		return nil, err
	}

	if found == nil {
		return nil, &entities.ErrEventDateNotFound{Date: created}
	}

	return found, nil
}

func (c *Calendar) Add(ctx context.Context, event *entities.Event) error {
	found, err := c.storage.Get(ctx, event.Id)
	if err != nil {
		return err
	}
	if found != nil {
		return &entities.ErrEventAlreadyExists{Id: event.Id}
	}

	found, _ = c.storage.Search(ctx, event.Created)
	if found != nil {
		return &entities.ErrDateBusy{Date: event.Created}
	}

	err = c.storage.Add(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Update(ctx context.Context, event *entities.Event) error {
	found, err := c.storage.Search(ctx, event.Created)
	if err != nil {
		return err
	}
	if found != nil && event.Id != found.Id {
		return &entities.ErrDateBusy{Date: event.Created}
	}

	current, _ := c.storage.Get(ctx, event.Id)
	if current == nil {
		return &entities.ErrEventNotFound{Id: event.Id}
	}

	err = c.storage.Update(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Delete(ctx context.Context, eventId int64) error {
	event, err := c.storage.Get(ctx, eventId)
	if err != nil {
		return err
	}
	if event == nil {
		return &entities.ErrEventNotFound{Id: eventId}
	}

	err = c.storage.Delete(ctx, eventId)
	if err != nil {
		return err
	}

	return nil
}
