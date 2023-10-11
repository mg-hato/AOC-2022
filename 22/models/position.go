package models

import c "aoc/common"

type Position = c.Pair[int, int]

func MakePosition(row, column int) Position {
	return c.MakePair(row, column)
}
