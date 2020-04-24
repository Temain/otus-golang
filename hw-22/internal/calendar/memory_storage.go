package calendar

import (
	"time"

	e "github.com/Temain/otus-golang/hw-22/internal/calendar/entities"
	i "github.com/Temain/otus-golang/hw-22/internal/calendar/interfaces"
)

type MemoryStorage struct {
	events map[int64]*e.Event
}

func NewMemoryStorage() i.ICalendarStorage {
	return &MemoryStorage{events: map[int64]*e.Event{}}
}
func CreateMemoryStorage(events map[int64]*e.Event) i.ICalendarStorage {
	return &MemoryStorage{events: events}
}

func (mc *MemoryStorage) List() ([]e.Event, error) {
	list := make([]e.Event, 0, len(mc.events))
	for _, event := range mc.events {
		list = append(list, *event)
	}
	return list, nil
}

func (mc *MemoryStorage) Search(created time.Time) (*e.Event, error) {
	for _, event := range mc.events {
		if event.Created == created {
			return event, nil
		}
	}

	return nil, nil
}

func (mc *MemoryStorage) Get(id int64) (*e.Event, error) {
	found, ok := mc.events[id]
	if !ok {
		return nil, &e.ErrEventNotFound{Id: id}
	}

	return found, nil
}

func (mc *MemoryStorage) Add(event *e.Event) error {
	mc.events[event.Id] = event

	return nil
}

func (mc *MemoryStorage) Update(event *e.Event) error {
	current := mc.events[event.Id]
	current.Title = event.Title
	current.Description = event.Description
	current.Created = event.Created

	return nil
}

func (mc *MemoryStorage) Delete(eventId int64) error {
	delete(mc.events, eventId)

	return nil
}
