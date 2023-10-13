package models

import c "aoc/common"

type Position c.Pair[int, int]

func MakePosition(row, col int) Position {
	return Position{
		First:  row,
		Second: col,
	}
}

func (p Position) Row() int {
	return p.First
}

func (p Position) Column() int {
	return p.Second
}

func (p Position) Add(other Position) Position {
	return MakePosition(
		p.Row()+other.Row(),
		p.Column()+other.Column(),
	)
}

func (p Position) Adjacent(d Direction) Position {
	row_delta, col_delta := 0, 0
	switch d {
	case North, NorthEast, NorthWest:
		row_delta = -1
	case South, SouthEast, SouthWest:
		row_delta = 1
	}

	switch d {
	case East, NorthEast, SouthEast:
		col_delta = 1
	case West, NorthWest, SouthWest:
		col_delta = -1
	}

	return MakePosition(p.Row()+row_delta, p.Column()+col_delta)
}
