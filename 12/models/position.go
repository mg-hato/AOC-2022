package models

import f "aoc/functional"

type Position = f.Pair[int, int]

// Get neighbours of the passed position
func GetNeighbours(position Position) []Position {
	return f.Map(func(direction f.Pair[int, int]) Position {
		return Position{
			First:  position.First + direction.First,
			Second: position.Second + direction.Second,
		}
	}, []f.Pair[int, int]{
		{First: 0, Second: 1},
		{First: 0, Second: -1},
		{First: 1, Second: 0},
		{First: -1, Second: 0},
	})
}
