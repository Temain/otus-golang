package entities

import "time"

type Statistics struct {
	TypeId   int64     `db:"type_id"`
	BannerId int64     `db:"banner_id"`
	SlotId   int64     `db:"slot_id"`
	GroupId  int64     `db:"group_id"`
	DateTime time.Time `db:"date_time"`
}
