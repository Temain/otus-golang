package algorithm

import (
	"fmt"
	"math"

	"github.com/Temain/otus-golang/project/pkg/algorithm/entities"
	"github.com/Temain/otus-golang/project/pkg/algorithm/interfaces"
)

type MultiarmedBandit struct {
}

func NewMultiarmedBandit() (interfaces.RotationAlgorithm, error) {
	return &MultiarmedBandit{}, nil
}

func (mb *MultiarmedBandit) GetHandle(items []entities.AlgorithmData) (int64, error) {
	if len(items) == 0 {
		return 0, fmt.Errorf("wrong data")
	}

	var countTotal int64
	for _, item := range items {
		if item.Count == 0 {
			return item.HandleId, nil
		}
		countTotal += item.Count
	}

	fmt.Printf("total count: %v\n", countTotal)

	var handleId int64
	var maxResult float64
	for _, item := range items {
		result := float64(item.AvgIncome) + math.Sqrt((2*math.Log(float64(countTotal)))/float64(item.Count))
		if result > maxResult {
			maxResult = result
			handleId = item.HandleId
		}

		fmt.Printf("handleId: %v, priority: %v\n", item.HandleId, result)
	}

	if math.IsInf(maxResult, 0) {
		return 0, fmt.Errorf("wrong result of algorthm, priority is infinity")
	}

	fmt.Printf("Handle with max priority: %v, priority: %v\n", handleId, maxResult)

	return handleId, nil
}
