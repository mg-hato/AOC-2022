package models

import c "aoc/common"

type valves_envelope struct {
	valves []Valve
}

func ValveEnvelope(valves ...Valve) c.Envelope[[]Valve] {
	return valves_envelope{valves}
}

func (ve valves_envelope) Get() []Valve {
	return c.Map(
		func(valve Valve) Valve {
			return Valve{
				valve.ID,
				valve.Flow_rate,
				c.ShallowCopy(valve.Tunnels),
			}
		},
		ve.valves,
	)
}

func ValvesEnvelopeEqualityFunc(lhs, rhs c.Envelope[[]Valve]) bool {
	return c.ArrayEqualWith(ValveEqualityFunc)(lhs.Get(), rhs.Get())
}
