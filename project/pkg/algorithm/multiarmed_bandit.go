package algorithm

import (
	"github.com/Temain/otus-golang/project/pkg/algorithm/entities"
	"github.com/Temain/otus-golang/project/pkg/algorithm/interfaces"
)

type MultiarmedBandit struct {
}

func NewMultiarmedBandit() (interfaces.RotationAlgorithm, error) {
	return &MultiarmedBandit{}, nil
}

func (mb *MultiarmedBandit) GetHandle([]entities.AlgorithmData) (int64, error) {
	return 0, nil
}
