package list

import "testing"

func TestItemValue(t *testing.T) {
	val := 1
	item := Item{val: val}
	if item.Value() != val {
		t.Fatalf("bad item value %d: expected %d", item.Value(), val)
	}
}

func TestItemPrev(t *testing.T) {
	item1 := Item{val: 1}
	item2 := Item{val: 2, prev: &item1}
	prev := item2.Prev()
	if prev.Value() != item1.Value() {
		t.Fatalf("bad item prev value %d: expected %d", prev.Value(), item1.Value())
	}
}

func TestItemNext(t *testing.T) {
	item2 := Item{val: 2}
	item1 := Item{val: 1, next: &item2}
	next := item1.Next()
	if next.Value() != item2.Value() {
		t.Fatalf("bad item next value %d: expected %d", next.Value(), item2.Value())
	}
}
