package models

import (
	c "aoc/common"
)

type terrain_envelope struct {
	terrain Terrain
}

func (te terrain_envelope) Get() Terrain {
	return c.Map(c.Identity[string], te.terrain)
}

func TerrainEnvelope(terrain ...string) c.Envelope[Terrain] {
	return terrain_envelope{terrain}
}

func TerrainEnvelopeEqualityFunction(lhs, rhs c.Envelope[Terrain]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
