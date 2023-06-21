package models

import (
	c "aoc/common"
	"fmt"
	"strings"
)

type MovingPlan struct {
	StartingContainers []Containers
	Moves              []Move
}

func (mp MovingPlan) String() string {
	return fmt.Sprintf(
		"Plan{containers: [%s]; %v}",
		strings.Join(
			c.Map(func(c Containers) string { return fmt.Sprintf(`"%s"`, c) }, mp.StartingContainers),
			", ",
		), mp.Moves,
	)
}

func MovingPlanEqualityFunction(lhs, rhs MovingPlan) bool {
	return c.ArrayEqual(lhs.StartingContainers, rhs.StartingContainers) &&
		c.ArrayEqual(lhs.Moves, rhs.Moves)
}
