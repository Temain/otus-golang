package interfaces

import "github.com/Temain/otus-golang/project/pkg/algorithm/entities"

type RotationAlgorithm interface {
	GetHandle([]entities.AlgorithmData) (int64, error)
}
