package models

import c "aoc/common"

type Valve struct {
	ID        string
	Flow_rate int
	Tunnels   []string
}

func ValveEqualityFunc(lhs, rhs Valve) bool {
	return lhs.ID == rhs.ID &&
		lhs.Flow_rate == rhs.Flow_rate &&
		c.ArrayEqual(lhs.Tunnels, rhs.Tunnels)
}
