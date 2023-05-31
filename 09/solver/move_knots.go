package solver

import m "aoc/day09/models"

func move_knots(direction m.Direction, knots []m.Position) {
	knots[0] = m.Move(knots[0], direction.AsMovement())
	for i := 1; i < len(knots); i++ {
		knots[i] = m.FollowLeader(knots[i-1], knots[i])
	}
}
