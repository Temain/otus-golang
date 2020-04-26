package calendar

import (
	"time"

	"github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
	interfaces "github.com/Temain/otus-golang/hw-21/internal/calendar/interfaces"
)

type Calendar struct {
	storage interfaces.EventStorage
}

func NewMemoryCalendar() interfaces.EventAdapter {
	storage := NewMemoryStorage()
	return &Calendar{storage}
}
func CreateCalendar(storage interfaces.EventStorage) interfaces.EventAdapter {
	return &Calendar{storage}
}

func (c *Calendar) List() ([]entities.Event, error) {
	list, err := c.storage.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Calendar) Search(created time.Time) (*entities.Event, error) {
	found, err := c.storage.Search(created)
	if err != nil {
		return nil, err
	}

	if found == nil {
		return nil, &entities.ErrEventDateNotFound{Date: created}
	}

	return found, nil
}

func (c *Calendar) Add(event *entities.Event) error {
	found, err := c.storage.Get(event.Id)
	if err != nil {
		return err
	}
	if found != nil {
		return &entities.ErrEventAlreadyExists{Id: event.Id}
	}

	found, _ = c.storage.Search(event.Created)
	if found != nil {
		return &entities.ErrDateBusy{Date: event.Created}
	}

	err = c.storage.Add(event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Update(event *entities.Event) error {
	found, err := c.storage.Search(event.Created)
	if err != nil {
		return err
	}
	if found != nil {
		return &entities.ErrDateBusy{Date: event.Created}
	}

	current, _ := c.storage.Get(event.Id)
	if current == nil {
		return &entities.ErrEventNotFound{Id: event.Id}
	}

	err = c.storage.Update(event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Delete(eventId int64) error {
	event, err := c.storage.Get(eventId)
	if err != nil {
		return err
	}
	if event == nil {
		return &entities.ErrEventNotFound{Id: eventId}
	}

	err = c.storage.Delete(eventId)
	if err != nil {
		return err
	}

	return nil
}
