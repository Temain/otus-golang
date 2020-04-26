package storages

import (
	"context"
	"sync"
	"time"

	"github.com/Temain/otus-golang/hw-28/internal/domain/entities"
	interfaces "github.com/Temain/otus-golang/hw-28/internal/domain/interfaces"
)

type MemoryStorage struct {
	index  int64
	events map[int64]*entities.Event
	mux    sync.RWMutex
}

func NewMemoryStorage() interfaces.EventStorage {
	return &MemoryStorage{index: 1, events: map[int64]*entities.Event{}}
}
func CreateMemoryStorage(events map[int64]*entities.Event) interfaces.EventStorage {
	return &MemoryStorage{index: 1, events: events}
}

func (mc *MemoryStorage) List(ctx context.Context) ([]entities.Event, error) {
	mc.mux.RLock()
	defer mc.mux.RUnlock()
	list := make([]entities.Event, 0, len(mc.events))
	for _, event := range mc.events {
		list = append(list, *event)
	}
	return list, nil
}

func (mc *MemoryStorage) Search(ctx context.Context, created time.Time) (*entities.Event, error) {
	mc.mux.RLock()
	defer mc.mux.RUnlock()
	for _, event := range mc.events {
		if event.Created == created {
			return event, nil
		}
	}

	return nil, nil
}

func (mc *MemoryStorage) Get(ctx context.Context, id int64) (*entities.Event, error) {
	mc.mux.RLock()
	defer mc.mux.RUnlock()
	found, ok := mc.events[id]
	if !ok {
		return nil, &entities.ErrEventNotFound{Id: id}
	}

	return found, nil
}

func (mc *MemoryStorage) Add(ctx context.Context, event *entities.Event) error {
	mc.mux.Lock()
	defer mc.mux.Unlock()
	event.Id = mc.index
	mc.events[event.Id] = event
	mc.index++

	return nil
}

func (mc *MemoryStorage) Update(ctx context.Context, event *entities.Event) error {
	mc.mux.Lock()
	defer mc.mux.Unlock()
	current := mc.events[event.Id]
	current.Title = event.Title
	current.Description = event.Description
	current.Created = event.Created

	return nil
}

func (mc *MemoryStorage) Delete(ctx context.Context, eventId int64) error {
	mc.mux.Lock()
	defer mc.mux.Unlock()
	delete(mc.events, eventId)

	return nil
}
