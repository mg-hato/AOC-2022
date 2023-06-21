package models

import (
	c "aoc/common"
)

func RockStructureEnvelope(rock_structures []RockStructure) c.Envelope[[]RockStructure] {
	return rock_structures_envelope{
		c.Filter(func(rs RockStructure) bool { return len(rs) > 0 }, rock_structures),
	}
}

type rock_structures_envelope struct {
	rock_structures []RockStructure
}

func (eee rock_structures_envelope) Get() []RockStructure {
	return c.Map(
		func(rs RockStructure) RockStructure { return c.Map(c.Identity[c.Pair[int, int]], rs) },
		eee.rock_structures,
	)
}

func RockStructureEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]RockStructure]) bool {
	return c.ArrayEqualWith(c.ArrayEqual[c.Pair[int, int]])(lhs.Get(), rhs.Get())
}
