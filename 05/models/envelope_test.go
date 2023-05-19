package models

import (
	"aoc/testers"
	"testing"
)

func TestD05_MovingPlanEnvelopeTest(t *testing.T) {
	data_generator := func() MovingPlan {
		return MovingPlan{
			[]string{"ABC", "DEF"},
			[]Move{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		}
	}
	envelope := CreateMovingPlanEnvelope(data_generator())
	plan := envelope.Get()
	plan.Moves[0].Destination = 10
	plan.StartingContainers[1] = "PQRST"
	testers.AssertEqualWithEqFunc(t, envelope.Get(), data_generator(), MovingPlanEqualityFunction)
}
