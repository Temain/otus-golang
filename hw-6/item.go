package hw_6

type Item struct {
}

func (i Item) Value() interface{} {
	return 0
}

func (i Item) Next() *Item {
	return &Item{}
}

func (i Item) Prev() *Item {
	return &Item{}
}
