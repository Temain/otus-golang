package entities

import "time"

type Rotation struct {
	BannerId  int64     `db:"banner_id"`
	SlotId    int64     `db:"slot_id"`
	StartedAt time.Time `db:"started_at"`
}
