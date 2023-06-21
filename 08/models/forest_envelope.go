package models

import (
	c "aoc/common"
)

type forest_envelope struct {
	forest Forest
}

func (env forest_envelope) Get() Forest {
	return c.Map(func(row []byte) []byte {
		return c.Map(c.Identity[byte], row)
	}, env.forest)
}

func ForestEnvelope(forest Forest) c.Envelope[Forest] {
	return forest_envelope{forest}
}

func ForestEnvelopeEqualityFunction(lhs, rhs c.Envelope[Forest]) bool {
	return c.ArrayEqualWith(c.ArrayEqual[byte])(lhs.Get(), rhs.Get())
}
