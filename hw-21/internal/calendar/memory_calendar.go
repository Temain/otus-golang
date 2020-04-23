package calendar

import (
	"time"

	e "github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
	i "github.com/Temain/otus-golang/hw-21/internal/calendar/interfaces"
)

type MemoryCalendar struct {
	events map[int64]*e.Event
}

func NewCalendar() i.Calendar {
	return &MemoryCalendar{events: map[int64]*e.Event{}}
}
func CreateCalendar(events map[int64]*e.Event) i.Calendar {
	return &MemoryCalendar{events: events}
}

func (mc *MemoryCalendar) List() ([]e.Event, error) {
	list := make([]e.Event, 0, len(mc.events))
	for _, event := range mc.events {
		list = append(list, *event)
	}
	return list, nil
}

func (mc *MemoryCalendar) Search(created time.Time) (*e.Event, error) {
	for _, event := range mc.events {
		if event.Created == created {
			return event, nil
		}
	}

	return nil, &e.ErrEventDateNotFound{Date: created}
}

func (mc *MemoryCalendar) Add(event *e.Event) error {
	found, ok := mc.events[event.Id]
	if ok {
		return &e.ErrEventAlreadyExists{Id: event.Id}
	}

	found, _ = mc.Search(event.Created)
	if found != nil {
		return &e.ErrDateBusy{Date: event.Created}
	}

	mc.events[event.Id] = event

	return nil
}

func (mc *MemoryCalendar) Update(event *e.Event) error {
	found, _ := mc.Search(event.Created)
	if found != nil {
		return &e.ErrDateBusy{Date: event.Created}
	}

	current, ok := mc.events[event.Id]
	if !ok {
		return &e.ErrEventNotFound{Id: event.Id}
	}

	current.Title = event.Title
	current.Description = event.Description
	current.Created = event.Created

	return nil
}

func (mc *MemoryCalendar) Delete(eventId int64) error {
	_, exists := mc.events[eventId]
	if !exists {
		return &e.ErrEventNotFound{Id: eventId}
	}

	delete(mc.events, eventId)
	_, exists = mc.events[eventId]
	if exists {
		return &e.ErrEventNotDeleted{Id: eventId}
	}

	return nil
}
