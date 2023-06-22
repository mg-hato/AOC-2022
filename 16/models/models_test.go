package models

import (
	c "aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestD16_EnvelopeTest(t *testing.T) {
	data := func() []Valve {
		return []Valve{
			{"AA", 0, []string{"BB", "CC"}},
			{"BB", 12, []string{"AA", "CC"}},
			{"CC", 7, []string{"AA", "BB"}},
		}
	}

	envelope := ValveEnvelope(data()...)

	valves := envelope.Get()
	valves[0].Flow_rate = 12
	valves[2].Tunnels[0] = "A1"

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data(), c.ArrayEqualWith(ValveEqualityFunc))
}
