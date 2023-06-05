package solver

import (
	m "aoc/d12/models"
	f "aoc/functional"
	"errors"
)

func StartingPositionDistancePicker(starting_position_letters ...rune) DistancePicker {
	return starting_position_distance_picker{
		f.CreateSet(starting_position_letters, f.Identity[rune]),
	}
}

type starting_position_distance_picker struct {
	starting_position_letters map[rune]bool
}

func (spdp starting_position_distance_picker) getDistance(distance_mappings map[m.Field]int) (int, error) {
	starting_positions := f.Filter(func(field m.Field) bool {
		return spdp.starting_position_letters[field.HeightCode]
	}, f.GetKeys(distance_mappings))

	if len(starting_positions) == 0 {
		return -1, errors.New("error while solving: the goal field cannot be reached from any starting position")
	}

	return f.Minimum(
		f.Map(func(field m.Field) int { return distance_mappings[field] }, starting_positions),
		func(lhs, rhs int) bool { return lhs < rhs },
	), nil
}
