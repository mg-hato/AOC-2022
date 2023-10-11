package models

import (
	c "aoc/common"
	"sort"
)

type Orientation int

const (
	North Orientation = 0
	East  Orientation = 1
	South Orientation = 2
	West  Orientation = 3
)

func GetOrientations() []Orientation {
	return []Orientation{North, East, South, West}
}

func (o Orientation) GetAdjacentOrientations() []Orientation {
	return []Orientation{
		Orientation((int(o) + 1) % 4),
		Orientation((int(o) + 3) % 4),
	}
}

func (orientation Orientation) String() string {
	switch orientation {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	default:
		return "[orientation:unknown]"
	}
}

func (o Orientation) GetFacingValue() int {
	return (int(o) + 3) % 4
}

func (o Orientation) move(pos Position, steps int) Position {
	row_adj, col_adj := 0, 0
	switch o {
	case North:
		row_adj = -steps
	case South:
		row_adj = steps
	case West:
		col_adj = -steps
	case East:
		col_adj = steps
	}

	return MakePosition(
		pos.First+row_adj,
		pos.Second+col_adj,
	)
}

// position orientation value: a number that describes given position w.r.t. orientation
// the bigger the number, the closer position is towards orientation `o`
func (o Orientation) pov(p Position) int {
	switch o {
	case North:
		return -p.First
	case South:
		return p.First
	case East:
		return p.Second
	case West:
		return -p.Second
	default:
		return 0
	}
}

// returns true iff lhs is more toward position `o` than rhs
func (o Orientation) compare(lhs, rhs Position) bool {
	return o.pov(lhs) > o.pov(rhs)
}

// Keeps positions that are facing most towards the given orientation (e.g. northmost)
func (o Orientation) most(ps []Position) []Position {
	if len(ps) == 0 {
		return []Position{}
	}
	best_pov := c.Maximum(c.Map(o.pov, ps))
	return c.Filter(func(p Position) bool { return o.pov(p) == best_pov }, ps)
}

// Sort positions based on which one is more towards given orientation `o`
func (o Orientation) order_towards(ps []Position) []Position {
	ps = c.ShallowCopy(ps)
	sort.Slice(ps, func(i, j int) bool {
		return o.pov(ps[i]) > o.pov(ps[j])
	})
	return ps
}

func (o Orientation) opposite() Orientation {
	return Orientation((int(o) + 2) % 4)
}

func get_orientation_rotator(base1, base2 Orientation) func(orientation_base1 Orientation) Orientation {
	rotation := int(base2) - int(base1)
	return func(ob1 Orientation) Orientation {

		return Orientation((int(ob1) + rotation + 4) % 4)
	}
}
