package models

import c "aoc/common"

type Direction string

const (
	UP    Direction = "U"
	DOWN  Direction = "D"
	LEFT  Direction = "L"
	RIGHT Direction = "R"
)

func (direction Direction) AsMovement() Movement {
	switch direction {
	case UP:
		return c.MakePair(0, 1)
	case DOWN:
		return c.MakePair(0, -1)
	case LEFT:
		return c.MakePair(-1, 0)
	case RIGHT:
		return c.MakePair(1, 0)
	default:
		return c.MakePair(0, 0)
	}
}
