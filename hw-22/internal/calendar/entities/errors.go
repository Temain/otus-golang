package entities

import (
	"fmt"
	"time"
)

type ErrDateBusy struct {
	Date time.Time
}

func (e *ErrDateBusy) Error() string {
	return fmt.Sprintf("date %v already busy", e.Date)
}

type ErrEventNotFound struct {
	Id int64
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
	Id int64
}

func (e *ErrEventAlreadyExists) Error() string {
	return fmt.Sprintf("event %d already exists", e.Id)
}

type ErrEventNotDeleted struct {
	Id int64
}

func (e *ErrEventNotDeleted) Error() string {
	return fmt.Sprintf("event %d not deleted", e.Id)
}
