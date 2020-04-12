package calendar

import (
	"fmt"
	"time"
)

type ErrDateBusy struct {
	date time.Time
}

func (e *ErrDateBusy) Error() string {
	return fmt.Sprintf("date %v already busy", e.date)
}

type ErrEventNotFound struct {
	Id int
}

func (e *ErrEventNotFound) Error() string {
	return fmt.Sprintf("event %d not found", e.Id)
}

type ErrEventDateNotFound struct {
	Date time.Time
}

func (e *ErrEventDateNotFound) Error() string {
	return fmt.Sprintf("event on date %v not found", e.Date)
}

type ErrEventAlreadyExists struct {
	Id int
}

func (e *ErrEventAlreadyExists) Error() string {
	return fmt.Sprintf("event %d already exists", e.Id)
}

type ErrEventNotDeleted struct {
	Id int
}

func (e *ErrEventNotDeleted) Error() string {
	return fmt.Sprintf("event %d not deleted", e.Id)
}
