package solver

import c "aoc/common"

type position = c.Pair[int, int]

func make_position(row, column int) position {
	return position{First: row, Second: column}
}
