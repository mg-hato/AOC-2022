package solver

import (
	m "aoc/d12/models"
)

type DistancePicker interface {
	getDistance(map[m.Field]int) (int, error)
}
