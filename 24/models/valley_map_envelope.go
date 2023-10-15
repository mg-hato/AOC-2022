package models

import c "aoc/common"

type valley_map_envelope struct {
	valley_map [][]rune
}

func CreateValleyMapEnvelope(valley_map [][]rune) c.Envelope[[][]rune] {
	return valley_map_envelope{valley_map}
}

func (vme valley_map_envelope) Get() [][]rune {
	return c.Map(c.ShallowCopy[rune], vme.valley_map)
}

func ValleyMapEnvelopeEqualityFunction(lhs, rhs c.Envelope[[][]rune]) bool {
	return c.ArrayEqualWith(c.ArrayEqual[rune])(lhs.Get(), rhs.Get())
}
