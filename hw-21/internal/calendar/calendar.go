package calendar

import (
	"time"
)

type Calendar interface {
	List() ([]Event, error)
	Search(date time.Time) (*Event, error)
	Add(*Event) error
	Update(*Event) error
	Delete(int) error
}

type MemoryCalendar struct {
	events map[int]*Event
}

func NewCalendar() Calendar {
	return &MemoryCalendar{events: map[int]*Event{}}
}
func CreateCalendar(events map[int]*Event) Calendar {
	return &MemoryCalendar{events: events}
}

func (c *MemoryCalendar) List() ([]Event, error) {
	list := make([]Event, 0, len(c.events))
	for _, event := range c.events {
		list = append(list, *event)
	}
	return list, nil
}

func (c *MemoryCalendar) Search(created time.Time) (*Event, error) {
	for _, event := range c.events {
		if event.Created == created {
			return event, nil
		}
	}

	return nil, &ErrEventDateNotFound{created}
}

func (c *MemoryCalendar) Add(event *Event) error {
	found, ok := c.events[event.Id]
	if ok {
		return &ErrEventAlreadyExists{event.Id}
	}

	found, _ = c.Search(event.Created)
	if found != nil {
		return &ErrDateBusy{event.Created}
	}

	c.events[event.Id] = event

	return nil
}

func (c *MemoryCalendar) Update(event *Event) error {
	found, _ := c.Search(event.Created)
	if found != nil {
		return &ErrDateBusy{event.Created}
	}

	current, ok := c.events[event.Id]
	if !ok {
		return &ErrEventNotFound{event.Id}
	}

	current.Title = event.Title
	current.Description = event.Description
	current.Created = event.Created

	return nil
}

func (c *MemoryCalendar) Delete(eventId int) error {
	delete(c.events, eventId)
	_, exists := c.events[eventId]
	if exists {
		return &ErrEventNotDeleted{eventId}
	}

	return nil
}
