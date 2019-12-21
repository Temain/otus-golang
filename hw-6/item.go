package hw_6

type Item struct {
	val  interface{}
	prev *Item
	next *Item
}

func (i Item) Value() interface{} {
	return i.val
}

func (i Item) Next() *Item {
	return i.next
}

func (i Item) Prev() *Item {
	return i.prev
}
