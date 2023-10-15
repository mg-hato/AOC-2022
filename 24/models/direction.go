package models

type Direction int

const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

func Directions() []Direction {
	return []Direction{North, East, South, West}
}

func (d Direction) String() string {
	switch d {
	case North:
		return "^"
	case East:
		return ">"
	case South:
		return "v"
	case West:
		return "<"
	default:
		return "<UNKNOWN_DIRECTION>"
	}
}

func (d Direction) Opposite() Direction {
	return Direction((int(d) + 2) % 4)
}
