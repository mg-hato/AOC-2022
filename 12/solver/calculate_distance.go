package solver

import (
	m "aoc/d12/models"
	e "aoc/envelope"
)

func CalculateDistance(distance_picker DistancePicker) func(e.Envelope[m.Terrain]) (int, error) {
	return func(envelope e.Envelope[m.Terrain]) (int, error) {
		return distance_picker.getDistance(findAllDistancesFromGoalPosition(m.EnumerateTerrain(envelope.Get())))
	}
}
