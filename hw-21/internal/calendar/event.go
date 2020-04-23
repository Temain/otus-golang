package calendar

import "time"

type Event struct {
	Id          int64
	Title       string
	Description string
	Created     time.Time
}
