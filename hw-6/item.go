package list

// Item элемент двусвязного списка.
type Item struct {
	val  interface{}
	prev *Item
	next *Item
}

// Value получить значение элемента.
func (i Item) Value() interface{} {
	return i.val
}

// Next получить указатель на следующий элемент списка.
func (i Item) Next() *Item {
	return i.next
}

// Prev получить указатель на предыдущий элемент списка.
func (i Item) Prev() *Item {
	return i.prev
}
