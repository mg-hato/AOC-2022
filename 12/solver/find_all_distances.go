package solver

import (
	m "aoc/d12/models"
	f "aoc/functional"
)

func findAllDistancesFromGoalPosition(terrain [][]m.Field) map[m.Field]int {

	// Find goal/end field and start from it backwards; if no such field return empty mapping
	goals := f.Filter(func(field m.Field) bool { return field.HeightCode == 'E' }, f.Flatten(terrain))
	if len(goals) == 0 {
		return map[m.Field]int{}
	}

	distances := map[m.Field]int{
		goals[0]: 0,
	}
	queue := make([]m.Field, 0)

	// a predicate function to test whether a position is within the terrain limits
	position_within_bounds := func(position m.Position) bool {
		return f.InRange(0, len(terrain))(position.First) && f.InRange(0, len(terrain[0]))(position.Second)
	}

	// (valid) position to terrain field transformation function
	position_to_field := func(position m.Position) m.Field { return terrain[position.First][position.Second] }

	// helper function that adds unexplored neighbouring terrain fields to the queue
	queue_up_neighbours := func(field m.Field) {
		queue = append(queue,
			f.Map(position_to_field, f.Filter(func(position m.Position) bool {
				if position_within_bounds(position) {
					_, already_explored := distances[terrain[position.First][position.Second]]
					return !already_explored
				}
				return false
			}, m.GetNeighbours(field.Position)),
			)...,
		)
	}

	queue_up_neighbours(goals[0])

	var i int = 0
	for i < len(queue) {
		field := queue[i]
		i++

		// ignore already explored fields
		if _, already_explored := distances[field]; already_explored {
			continue
		}

		// get already explored neighbours of the current field that are reachable in one step
		reachable_neighbours := f.Filter(
			func(neighbouring_field m.Field) bool {
				_, already_explored := distances[neighbouring_field]
				return already_explored && field.GetHeight()+1 >= neighbouring_field.GetHeight()
			},
			f.Map(position_to_field, f.Filter(position_within_bounds, m.GetNeighbours(field.Position))),
		)

		// If no such neighbours, ignore the current field for the time being
		if len(reachable_neighbours) == 0 {
			continue
		}

		// otherwise, calculate the shortest distance for the current field and queue up its neighbours for inspection
		closest_neighbour := f.Minimum(reachable_neighbours, func(lhs, rhs m.Field) bool {
			return distances[lhs] < distances[rhs]
		})
		distances[field] = 1 + distances[closest_neighbour]
		queue_up_neighbours(field)
	}
	return distances
}
