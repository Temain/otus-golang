package hw_6

type List struct {
	len   int
	first *Item
	last  *Item
}

func (list List) Len() int {
	return list.len
}

func (list List) First() *Item {
	return list.first
}

func (list List) Last() *Item {
	return list.last
}

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
