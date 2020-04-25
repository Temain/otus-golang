package calendar

import (
	"context"
	"time"

	e "github.com/Temain/otus-golang/hw-22/internal/calendar/entities"
	i "github.com/Temain/otus-golang/hw-22/internal/calendar/interfaces"
	s "github.com/Temain/otus-golang/hw-22/internal/calendar/storages"
)

type Calendar struct {
	storage i.EventStorage
}

func NewMemoryCalendar() i.EventAdapter {
	storage := s.NewMemoryStorage()
	return &Calendar{storage}
}
func NewPostgreCalendar(dsn string) (i.EventAdapter, error) {
	storage, err := s.NewPostgreStorage(dsn)
	if err != nil {
		return nil, err
	}
	return &Calendar{storage}, nil
}
func CreateCalendar(storage i.EventStorage) i.EventAdapter {
	return &Calendar{storage}
}

func (c *Calendar) List(ctx context.Context) ([]e.Event, error) {
	list, err := c.storage.List(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Calendar) Search(ctx context.Context, created time.Time) (*e.Event, error) {
	found, err := c.storage.Search(ctx, created)
	if err != nil {
		return nil, err
	}

	if found == nil {
		return nil, &e.ErrEventDateNotFound{Date: created}
	}

	return found, nil
}

func (c *Calendar) Add(ctx context.Context, event *e.Event) error {
	found, _ := c.storage.Get(ctx, event.Id)
	if found != nil {
		return &e.ErrEventAlreadyExists{Id: event.Id}
	}

	found, _ = c.storage.Search(ctx, event.Created)
	if found != nil {
		return &e.ErrDateBusy{Date: event.Created}
	}

	err := c.storage.Add(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Update(ctx context.Context, event *e.Event) error {
	found, _ := c.storage.Search(ctx, event.Created)
	if found != nil {
		return &e.ErrDateBusy{Date: event.Created}
	}

	current, _ := c.storage.Get(ctx, event.Id)
	if current == nil {
		return &e.ErrEventNotFound{Id: event.Id}
	}

	err := c.storage.Update(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Delete(ctx context.Context, eventId int64) error {
	event, _ := c.storage.Get(ctx, eventId)
	if event == nil {
		return &e.ErrEventNotFound{Id: eventId}
	}

	err := c.storage.Delete(ctx, eventId)
	if err != nil {
		return err
	}

	return nil
}
