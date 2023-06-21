package models

import c "aoc/common"

type Position = c.Pair[int, int]

// Get neighbours of the passed position
func GetNeighbours(position Position) []Position {
	return c.Map(func(direction c.Pair[int, int]) Position {
		return Position{
			First:  position.First + direction.First,
			Second: position.Second + direction.Second,
		}
	}, []c.Pair[int, int]{
		{First: 0, Second: 1},
		{First: 0, Second: -1},
		{First: 1, Second: 0},
		{First: -1, Second: 0},
	})
}
