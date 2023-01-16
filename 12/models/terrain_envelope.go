package models

import (
	e "aoc/envelope"
	f "aoc/functional"
)

type terrain_envelope struct {
	terrain Terrain
}

func (te terrain_envelope) Get() Terrain {
	return f.Map(f.Identity[string], te.terrain)
}

func TerrainEnvelope(terrain ...string) e.Envelope[Terrain] {
	return terrain_envelope{terrain}
}

func TerrainEnvelopeEqualityFunction(lhs, rhs e.Envelope[Terrain]) bool {
	return f.ArrayEqual(lhs.Get(), rhs.Get())
}
