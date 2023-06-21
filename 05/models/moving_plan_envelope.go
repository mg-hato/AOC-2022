package models

import (
	c "aoc/common"
	"fmt"
)

type MovingPlanEnvelope struct {
	plan MovingPlan
}

func (envelope MovingPlanEnvelope) String() string {
	return fmt.Sprintf("MovingPlanEnvelope[%s]", envelope.plan)
}

func (envelope MovingPlanEnvelope) Get() MovingPlan {
	data := MovingPlan{
		StartingContainers: make([]Containers, len(envelope.plan.StartingContainers)),
		Moves:              make([]Move, len(envelope.plan.Moves)),
	}
	copy(data.StartingContainers, envelope.plan.StartingContainers)
	copy(data.Moves, envelope.plan.Moves)
	return data
}

func CreateMovingPlanEnvelope(plan MovingPlan) c.Envelope[MovingPlan] {
	return &MovingPlanEnvelope{plan}
}

func MovingPlanEnvelopeEqualityFunction(lhs, rhs c.Envelope[MovingPlan]) bool {
	return MovingPlanEqualityFunction(lhs.Get(), rhs.Get())
}
