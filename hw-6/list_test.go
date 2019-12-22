package list

import (
	"testing"
)

func prepareList() List {
	list := List{}
	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)
	return list
}

func TestListLen(t *testing.T) {
	list := prepareList()
	expLen := 3
	listLen := list.Len()
	if listLen != expLen {
		t.Fatalf("bad list len %d: expected %d", listLen, expLen)
	}
}

func TestListFirst(t *testing.T) {
	list := prepareList()
	expFirst := 1
	listFirst := list.First()
	if listFirst.Value() != expFirst {
		t.Fatalf("bad first item %d: expected %d", listFirst.Value(), expFirst)
	}
}

func TestListLast(t *testing.T) {
	list := prepareList()
	expLast := 3
	listLast := list.Last()
	if listLast.Value() != expLast {
		t.Fatalf("bad last item %d: expected %d", listLast.Value(), expLast)
	}
}

func TestListPushFront(t *testing.T) {
	list := prepareList()
	newVal := 0
	list.PushFront(newVal)
	listFirst := list.First()
	if listFirst.Value() != newVal {
		t.Fatalf("bad first item %d after push front: expected %d", listFirst.Value(), newVal)
	}

	firstNext := listFirst.Next()
	if firstNext.Value() != newVal+1 {
		t.Fatalf("bad first next item %d after push front: expected %d", firstNext.Value(), newVal+1)
	}
}

func TestListPushBack(t *testing.T) {
	list := prepareList()
	newVal := 4
	list.PushBack(newVal)
	listLast := list.Last()
	if listLast.Value() != newVal {
		t.Fatalf("bad last item %d after push back: expected %d", listLast.Value(), newVal)
	}

	lastPrev := listLast.Prev()
	if lastPrev.Value() != newVal-1 {
		t.Fatalf("bad last prev item %d after push back: expected %d", lastPrev.Value(), newVal-1)
	}
}

func TestListRemove(t *testing.T) {
	list := prepareList()
	first := list.First()
	list.Remove(*first)
	newFirst := list.First()
	expLen := 2
	if list.len != expLen {
		t.Fatalf("bad list len %d after remove: expected %d", list.len, expLen)
	}

	expVal := 2
	if newFirst.Value() != expVal {
		t.Fatalf("bad first item %d after remove: expected %d", newFirst.Value(), expVal)
	}
}
