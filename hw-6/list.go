package list

// List двусвязный список.
type List struct {
	len   int
	first *Item
	last  *Item
}

// Len получить длинну списка.
func (list List) Len() int {
	return list.len
}

// First получить указатель на первый элемент списка.
func (list List) First() *Item {
	return list.first
}

// Last получить указатель на последний элемент списка.
func (list List) Last() *Item {
	return list.last
}

// PushFront добавить элемент со значнием v в начало списка.
func (list *List) PushFront(v interface{}) {
	item := Item{val: v}
	temp := list.first
	item.next = temp
	list.first = &item
	if list.len == 0 {
		list.last = list.first
	} else {
		temp.prev = &item
	}
	list.len++
}

// PushBack добавить элемент со значнием v в конец списка.
func (list *List) PushBack(v interface{}) {
	item := Item{val: v}
	temp := list.last
	item.prev = temp
	list.last = &item
	if list.len == 0 {
		list.first = list.last
	} else {
		temp.next = &item
	}
	list.len++
}

// Remove удалить элемент из списка.
func (list *List) Remove(item Item) {
	prev := item.prev
	next := item.next
	if prev != nil {
		prev.next = next
	} else {
		list.first = next
	}
	if next != nil {
		next.prev = prev
	} else {
		list.last = prev
	}
	list.len--
	item.prev = nil
	item.next = nil
}
