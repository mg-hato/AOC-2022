package models

import (
	c "aoc/common"
	"sort"
)

type Movement = c.Pair[int, int]
type Position = c.Pair[int, int]

func Move(position Position, movement Movement) Position {
	return c.MakePair(
		position.First+movement.First,
		position.Second+movement.Second,
	)
}

func CombineMovements(lhs, rhs Movement) Movement {
	return Move(lhs, rhs)
}

func absolute(value int) int {
	if value >= 0 {
		return value
	} else {
		return -value
	}
}

// More specifically: Manhattan distance
func Distance(x, y Position) int {
	return absolute(x.First-y.First) + absolute(x.Second-y.Second)
}

// Returns true if and only if provided positions are touching
//
// i.e. maximum difference in each axis (First and Second) is 1
func AreTouching(x, y Position) bool {
	return c.InInclusiveRange(-1, 1)(x.First-y.First) && c.InInclusiveRange(-1, 1)(x.Second-y.Second)
}

// Returns follower's new position
func FollowLeader(leader, follower Position) Position {

	if AreTouching(leader, follower) {
		return follower
	}

	new_position_candidates := c.Map(
		func(movement Movement) Position { return Move(follower, movement) },
		[]Movement{
			UP.AsMovement(), DOWN.AsMovement(),
			LEFT.AsMovement(), RIGHT.AsMovement(),
			CombineMovements(UP.AsMovement(), LEFT.AsMovement()),
			CombineMovements(UP.AsMovement(), RIGHT.AsMovement()),
			CombineMovements(DOWN.AsMovement(), LEFT.AsMovement()),
			CombineMovements(DOWN.AsMovement(), RIGHT.AsMovement()),
		},
	)

	sort.Slice(new_position_candidates, func(i, j int) bool {
		return Distance(leader, new_position_candidates[i]) < Distance(leader, new_position_candidates[j])
	})
	return new_position_candidates[0]
}
