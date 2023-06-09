package models

import (
	e "aoc/envelope"
	f "aoc/functional"
)

func RockStructureEnvelope(rock_structures []RockStructure) e.Envelope[[]RockStructure] {
	return rock_structures_envelope{
		f.Filter(func(rs RockStructure) bool { return len(rs) > 0 }, rock_structures),
	}
}

type rock_structures_envelope struct {
	rock_structures []RockStructure
}

func (eee rock_structures_envelope) Get() []RockStructure {
	return f.Map(
		func(rs RockStructure) RockStructure { return f.Map(f.Identity[f.Pair[int, int]], rs) },
		eee.rock_structures,
	)
}

func RockStructureEnvelopeEqualityFunction(lhs, rhs e.Envelope[[]RockStructure]) bool {
	return f.ArrayEqualWith(f.ArrayEqual[f.Pair[int, int]])(lhs.Get(), rhs.Get())
}
