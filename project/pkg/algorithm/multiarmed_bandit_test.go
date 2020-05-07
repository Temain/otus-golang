package algorithm

import (
	"testing"

	"github.com/Temain/otus-golang/project/pkg/algorithm/entities"
)

func TestGetHandle(t *testing.T) {
	algorithm, err := NewMultiarmedBandit()
	if err != nil {
		t.Fatalf("error on init algorithm: %v", err)
	}

	data := []entities.AlgorithmData{
		{HandleId: 1, Count: 10, AvgIncome: 3},
		{HandleId: 2, Count: 15, AvgIncome: 1},
		{HandleId: 3, Count: 3, AvgIncome: 0},
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
