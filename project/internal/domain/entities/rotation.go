package entities

import "time"

type Rotation struct {
	BannerId  int64
	SlotId    int64
	StartedAt time.Time
}
