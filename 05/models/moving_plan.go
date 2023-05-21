package models

import (
	f "aoc/functional"
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
			f.Map(func(c Containers) string { return fmt.Sprintf(`"%s"`, c) }, mp.StartingContainers),
			", ",
		), mp.Moves,
	)
}

func MovingPlanEqualityFunction(lhs, rhs MovingPlan) bool {
	return f.ArrayEqual(lhs.StartingContainers, rhs.StartingContainers) &&
		f.ArrayEqual(lhs.Moves, rhs.Moves)
}
