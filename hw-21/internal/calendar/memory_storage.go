package calendar

import (
	"time"

	"github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
	interfaces "github.com/Temain/otus-golang/hw-21/internal/calendar/interfaces"
)

var index int64 = 1

type MemoryStorage struct {
	events map[int64]*entities.Event
}

func NewMemoryStorage() interfaces.EventStorage {
	return &MemoryStorage{events: map[int64]*entities.Event{}}
}
func CreateMemoryStorage(events map[int64]*entities.Event) interfaces.EventStorage {
	return &MemoryStorage{events: events}
}

func (mc *MemoryStorage) List() ([]entities.Event, error) {
	list := make([]entities.Event, 0, len(mc.events))
	for _, event := range mc.events {
		list = append(list, *event)
	}
	return list, nil
}

func (mc *MemoryStorage) Search(created time.Time) (*entities.Event, error) {
	for _, event := range mc.events {
		if event.Created == created {
			return event, nil
		}
	}

	return nil, nil
}

func (mc *MemoryStorage) Get(id int64) (*entities.Event, error) {
	found, ok := mc.events[id]
	if !ok {
		return nil, &entities.ErrEventNotFound{Id: id}
	}

	return found, nil
}

func (mc *MemoryStorage) Add(event *entities.Event) error {
	event.Id = index
	mc.events[event.Id] = event
	index++

	return nil
}

func (mc *MemoryStorage) Update(event *entities.Event) error {
	current := mc.events[event.Id]
	current.Title = event.Title
	current.Description = event.Description
	current.Created = event.Created

	return nil
}

func (mc *MemoryStorage) Delete(id int64) error {
	delete(mc.events, id)

	return nil
}
