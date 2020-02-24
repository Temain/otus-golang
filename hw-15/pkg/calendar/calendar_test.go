package calendar

import (
	"testing"
	"time"
)

func TestAddEvent(t *testing.T) {
	c := NewCalendar()
	event := &Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Date:        time.Now(),
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
	c := NewCalendar()
	event := &Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Date:        time.Now(),
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
	c := NewCalendar()
	events, err := c.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 0 {
		t.Fatal("bad result of list, should be empty list")
	}

	event := &Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Date:        time.Now(),
	}
	c = CreateCalendar(map[int]*Event{1: event})
	events, err = c.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 {
		t.Fatal("bad result of list, expected 1 event")
	}
}

func TestSearchEvent(t *testing.T) {
	date := time.Now()
	event := &Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Date:        date,
	}
	c := CreateCalendar(map[int]*Event{1: event})
	found, _ := c.Search(date)
	if found == nil {
		t.Fatalf("bad search result, event not found")
	}
	if found.Date != date {
		t.Fatalf("bad search result, expected event date %v", date)
	}
}

func TestUpdateEvent(t *testing.T) {
	date := time.Now()
	event := &Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Date:        date,
	}
	c := CreateCalendar(map[int]*Event{1: event})
	eventNew := &Event{
		Id:          1,
		Title:       "Evening tea",
		Description: "Not bad",
		Date:        date.Add(time.Second),
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
		t.Fatalf("bad update result, expected event title %v", date)
	}
	if updated.Description != eventNew.Description {
		t.Fatalf("bad update result, expected event description %v", date)
	}
	if updated.Date != eventNew.Date {
		t.Fatalf("bad update result, expected event date %v", date)
	}
}

func TestDeleteEvent(t *testing.T) {
	date := time.Now()
	event := &Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Date:        date,
	}
	c := CreateCalendar(map[int]*Event{1: event})
	events, _ := c.List()
	if len(events) == 0 {
		t.Fatalf("bad result, prepared caledar is empty")
	}

	err := c.Delete(event.Id)
	if err != nil {
		t.Fatalf("bad delete result, %v", err)
	}
}
