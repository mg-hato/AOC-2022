package models

import c "aoc/common"

type blueprints_envelope struct {
	blueprints []Blueprint
}

func BlueprintsEnvelope(blueprints ...Blueprint) c.Envelope[[]Blueprint] {
	return blueprints_envelope{blueprints: blueprints}
}

func (envelope blueprints_envelope) Get() []Blueprint {
	return c.ShallowCopy(envelope.blueprints)
}

func BlueprintsEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]Blueprint]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
