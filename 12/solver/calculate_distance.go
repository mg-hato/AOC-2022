package solver

import (
	m "aoc/d12/models"
)

func CalculateDistance(distance_picker DistancePicker) func(m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {
		return distance_picker.getDistance(findAllDistancesFromGoalPosition(m.EnumerateTerrain(input.Get())))
	}
}
