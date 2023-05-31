package models

import f "aoc/functional"

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
		return f.MakePair(0, 1)
	case DOWN:
		return f.MakePair(0, -1)
	case LEFT:
		return f.MakePair(-1, 0)
	case RIGHT:
		return f.MakePair(1, 0)
	default:
		return f.MakePair(0, 0)
	}
}
