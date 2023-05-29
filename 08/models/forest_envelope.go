package models

import (
	e "aoc/envelope"
	f "aoc/functional"
)

type forest_envelope struct {
	forest Forest
}

func (env forest_envelope) Get() Forest {
	return f.Map(func(row []byte) []byte {
		return f.Map(f.Identity[byte], row)
	}, env.forest)
}

func ForestEnvelope(forest Forest) e.Envelope[Forest] {
	return forest_envelope{forest}
}

func ForestEnvelopeEqualityFunction(lhs, rhs e.Envelope[Forest]) bool {
	return f.ArrayEqualWith(f.ArrayEqual[byte])(lhs.Get(), rhs.Get())
}
