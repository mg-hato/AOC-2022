package models

import c "aoc/common"

type droplets_envelope struct {
	droplets []Droplet
}

func DropletsEnvelope(droplets ...Droplet) c.Envelope[[]Droplet] {
	return droplets_envelope{droplets}
}

func (envelope droplets_envelope) Get() []Droplet {
	return c.ShallowCopy(envelope.droplets)
}

func DropletsEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]Droplet]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
