package entities

import "time"

type Rotation struct {
	BannerId  int
	SlotId    int
	StartedAt time.Time
}
