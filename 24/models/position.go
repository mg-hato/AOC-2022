package models

import "aoc/common"

type Position common.Pair[int, int]

func MakePosition(row, col int) Position {
	return Position{First: row, Second: col}
}

func (p Position) Row() int { return p.First }

func (p Position) Column() int { return p.Second }

func (p Position) Move(d Direction) Position {
	row_delta, col_delta := 0, 0
	switch d {
	case North:
		row_delta = -1
	case South:
		row_delta = 1
	case West:
		col_delta = -1
	case East:
		col_delta = 1
	}
	return MakePosition(p.Row()+row_delta, p.Column()+col_delta)
}
