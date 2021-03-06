package algorithm

import (
	"testing"

	"github.com/Temain/otus-golang/project/algorithm/entities"
)

func TestGetHandle(t *testing.T) {
	algorithm, err := NewMultiarmedBandit()
	if err != nil {
		t.Fatalf("error on init algorithm: %v", err)
	}

	data := []entities.AlgorithmData{
		{HandleId: 1, Count: 1000, AvgIncome: 1},
		{HandleId: 2, Count: 1, AvgIncome: 0},
		{HandleId: 3, Count: 1, AvgIncome: 0},
	}
	handleId, err := algorithm.GetHandle(data)
	if err != nil {
		t.Fatalf("bad result, %v", err)
	}
	if handleId == 0 {
		t.Fatalf("bad result, handle id can't be 0")
	}
}

func TestGetHandleInit(t *testing.T) {
	algorithm, err := NewMultiarmedBandit()
	if err != nil {
		t.Fatalf("error on init algorithm: %v", err)
	}

	data := []entities.AlgorithmData{
		{HandleId: 1, Count: 0, AvgIncome: 0},
		{HandleId: 2, Count: 0, AvgIncome: 0},
		{HandleId: 3, Count: 0, AvgIncome: 0},
	}
	handleId, err := algorithm.GetHandle(data)
	if err != nil {
		t.Fatalf("bad result, %v", err)
	}
	if handleId == 0 {
		t.Fatalf("bad result, handle id can't be 0")
	}
}

func TestGetHandleNoItems(t *testing.T) {
	algorithm, err := NewMultiarmedBandit()
	if err != nil {
		t.Fatalf("error on init algorithm: %v", err)
	}

	var data []entities.AlgorithmData
	handleId, err := algorithm.GetHandle(data)
	if err == nil {
		t.Fatalf("bad result, with no data shound be error")
	}
	if handleId != 0 {
		t.Fatalf("bad result, handle id should be 0")
	}
}
