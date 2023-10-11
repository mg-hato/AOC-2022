package models

import c "aoc/common"

type grove_map_envelope struct {
	grove_map [][]SpotType
}

func CreateGroveMapEnvelope(grove_map [][]SpotType) c.Envelope[[][]SpotType] {
	return grove_map_envelope{grove_map}
}

func (gme grove_map_envelope) Get() [][]SpotType {
	return c.Map(c.ShallowCopy[SpotType], gme.grove_map)
}

func GroveMapEnvelopeEqualityFunc(lhs, rhs c.Envelope[[][]SpotType]) bool {
	return c.ArrayEqualWith(c.ArrayEqual[SpotType])(lhs.Get(), rhs.Get())
}
