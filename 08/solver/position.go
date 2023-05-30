package solver

import "aoc/functional"

type position = functional.Pair[int, int]

func make_position(row, column int) position {
	return position{First: row, Second: column}
}
