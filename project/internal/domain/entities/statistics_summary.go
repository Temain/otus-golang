package entities

type StatisticsSummary struct {
	BannerId int64 `db:"banner_id"`
	SlotId   int64 `db:"slot_id"`
	GroupId  int64 `db:"group_id"`
	Buyouts  int64 `db:"buyouts"`
	Clicks   int64 `db:"clicks"`
}
