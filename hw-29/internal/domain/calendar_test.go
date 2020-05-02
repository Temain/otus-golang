package domain

import (
	"context"
	"testing"
	"time"

	"github.com/Temain/otus-golang/hw-29/internal/domain/entities"
	"github.com/Temain/otus-golang/hw-29/internal/domain/storages"
)

func TestAddEvent(t *testing.T) {
	ctx := context.Background()
	c := NewMemoryCalendar()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     time.Now(),
	}
	err := c.Add(ctx, event)
	if err != nil {
		t.Fatalf("bad result on add event, %v", err)
	}

	events, err := c.List(ctx)
	if err != nil {
		t.Fatalf("bad result on list events, %v", err)
	}
	count := len(events)
	if count != 1 {
		t.Fatalf("bad events count after add: %d, expected: %d", count, 1)
	}
}

func TestAddDuplicateEvent(t *testing.T) {
	ctx := context.Background()
	c := NewMemoryCalendar()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     time.Now(),
	}
	err := c.Add(ctx, event)
	if err != nil {
		t.Fatalf("bad result on add event, %v", err)
	}
	err = c.Add(ctx, event)
	if err == nil {
		t.Fatalf("bad result on second add event, should be error")
	}
}

func TestList(t *testing.T) {
	ctx := context.Background()
	c := NewMemoryCalendar()
	events, err := c.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 0 {
		t.Fatal("bad result of list, should be empty list")
	}

	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     time.Now(),
	}
	s := storages.CreateMemoryStorage(map[int64]*entities.Event{1: event})
	c = CreateCalendar(s)
	events, err = c.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 {
		t.Fatal("bad result of list, expected 1 event")
	}
}

func TestSearchEvent(t *testing.T) {
	ctx := context.Background()
	created := time.Now()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     created,
	}
	s := storages.CreateMemoryStorage(map[int64]*entities.Event{1: event})
	c := CreateCalendar(s)
	found, _ := c.Search(ctx, created)
	if found == nil {
		t.Fatalf("bad search result, event not found")
	}
	if found.Created != created {
		t.Fatalf("bad search result, expected event date %v", created)
	}
}

func TestUpdateEvent(t *testing.T) {
	ctx := context.Background()
	created := time.Now()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     created,
	}
	s := storages.CreateMemoryStorage(map[int64]*entities.Event{1: event})
	c := CreateCalendar(s)
	eventNew := &entities.Event{
		Id:          1,
		Title:       "Evening tea",
		Description: "Not bad",
		Created:     created.Add(time.Second),
	}
	err := c.Update(ctx, eventNew)
	if err != nil {
		t.Fatalf("bad update result, %v", err)
	}
	events, _ := c.List(ctx)
	if len(events) == 0 {

	}
	updated := events[0]
	if updated.Title != eventNew.Title {
		t.Fatalf("bad update result, expected event title %v", created)
	}
	if updated.Description != eventNew.Description {
		t.Fatalf("bad update result, expected event description %v", created)
	}
	if updated.Created != eventNew.Created {
		t.Fatalf("bad update result, expected event date %v", created)
	}
}

func TestDeleteEvent(t *testing.T) {
	ctx := context.Background()
	created := time.Now()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     created,
	}
	s := storages.CreateMemoryStorage(map[int64]*entities.Event{1: event})
	c := CreateCalendar(s)
	events, _ := c.List(ctx)
	if len(events) == 0 {
		t.Fatalf("bad result, prepared caledar is empty")
	}

	err := c.Delete(ctx, event.Id)
	if err != nil {
		t.Fatalf("bad delete result, %v", err)
	}
}
