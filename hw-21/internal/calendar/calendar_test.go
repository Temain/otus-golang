package calendar

import (
	"testing"
	"time"

	"github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
)

func TestAddEvent(t *testing.T) {
	c := NewMemoryCalendar()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     time.Now(),
	}
	err := c.Add(event)
	if err != nil {
		t.Fatalf("bad result on add event, %v", err)
	}

	events, err := c.List()
	if err != nil {
		t.Fatalf("bad result on list events, %v", err)
	}
	count := len(events)
	if count != 1 {
		t.Fatalf("bad events count after add: %d, expected: %d", count, 1)
	}
}

func TestAddDuplicateEvent(t *testing.T) {
	c := NewMemoryCalendar()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     time.Now(),
	}
	err := c.Add(event)
	if err != nil {
		t.Fatalf("bad result on add event, %v", err)
	}
	err = c.Add(event)
	if err == nil {
		t.Fatalf("bad result on second add event, should be error")
	}
}

func TestList(t *testing.T) {
	c := NewMemoryCalendar()
	events, err := c.List()
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
	s := CreateMemoryStorage(map[int64]*entities.Event{1: event})
	c = CreateCalendar(s)
	events, err = c.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 {
		t.Fatal("bad result of list, expected 1 event")
	}
}

func TestSearchEvent(t *testing.T) {
	created := time.Now()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     created,
	}
	s := CreateMemoryStorage(map[int64]*entities.Event{1: event})
	c := CreateCalendar(s)
	found, _ := c.Search(created)
	if found == nil {
		t.Fatalf("bad search result, event not found")
	}
	if found.Created != created {
		t.Fatalf("bad search result, expected event date %v", created)
	}
}

func TestUpdateEvent(t *testing.T) {
	created := time.Now()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     created,
	}
	s := CreateMemoryStorage(map[int64]*entities.Event{1: event})
	c := CreateCalendar(s)
	eventNew := &entities.Event{
		Id:          1,
		Title:       "Evening tea",
		Description: "Not bad",
		Created:     created.Add(time.Second),
	}
	err := c.Update(eventNew)
	if err != nil {
		t.Fatalf("bad update result, %v", err)
	}
	events, _ := c.List()
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
	created := time.Now()
	event := &entities.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     created,
	}
	s := CreateMemoryStorage(map[int64]*entities.Event{1: event})
	c := CreateCalendar(s)
	events, _ := c.List()
	if len(events) == 0 {
		t.Fatalf("bad result, prepared caledar is empty")
	}

	err := c.Delete(event.Id)
	if err != nil {
		t.Fatalf("bad delete result, %v", err)
	}
}
