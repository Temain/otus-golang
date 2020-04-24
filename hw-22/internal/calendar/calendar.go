package calendar

import (
	"time"

	e "github.com/Temain/otus-golang/hw-22/internal/calendar/entities"
	i "github.com/Temain/otus-golang/hw-22/internal/calendar/interfaces"
)

type Calendar struct {
	storage i.ICalendarStorage
}

func NewMemoryCalendar() i.ICalendar {
	storage := NewMemoryStorage()
	return &Calendar{storage}
}
func CreateCalendar(storage i.ICalendarStorage) i.ICalendar {
	return &Calendar{storage}
}

func (c *Calendar) List() ([]e.Event, error) {
	list, err := c.storage.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Calendar) Search(created time.Time) (*e.Event, error) {
	found, err := c.storage.Search(created)
	if err != nil {
		return nil, err
	}

	if found == nil {
		return nil, &e.ErrEventDateNotFound{Date: created}
	}

	return found, nil
}

func (c *Calendar) Add(event *e.Event) error {
	found, _ := c.storage.Get(event.Id)
	if found != nil {
		return &e.ErrEventAlreadyExists{Id: event.Id}
	}

	found, _ = c.storage.Search(event.Created)
	if found != nil {
		return &e.ErrDateBusy{Date: event.Created}
	}

	err := c.storage.Add(event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Update(event *e.Event) error {
	found, _ := c.storage.Search(event.Created)
	if found != nil {
		return &e.ErrDateBusy{Date: event.Created}
	}

	current, _ := c.storage.Get(event.Id)
	if current == nil {
		return &e.ErrEventNotFound{Id: event.Id}
	}

	err := c.storage.Update(event)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Delete(eventId int64) error {
	event, _ := c.storage.Get(eventId)
	if event == nil {
		return &e.ErrEventNotFound{Id: eventId}
	}

	err := c.storage.Delete(eventId)
	if err != nil {
		return err
	}

	return nil
}
