package solver

import (
	c "aoc/common"
	m "aoc/d05/models"
	"strings"
)

func FollowMovingPlanWith(crate_mover CrateMover) func(m.SolverInput) (string, error) {
	return func(input m.SolverInput) (string, error) {
		moving_plan := input.Get()
		stacks := moving_plan.StartingContainers

		// Execute all the move instructions in order
		for i, move := range moving_plan.Moves {
			if !crate_mover.ExecuteMove(stacks, move) {
				return "", not_enough_containers_error(i+1, move)
			}
		}

		// Once finished, ensure that each stack is non-empty
		if c.Any(func(stack m.Containers) bool { return len(stack) == 0 }, stacks) {
			return "", empty_stacks_error(stacks)
		}

		// All ok, return result
		top_containers := strings.Join(c.Map(
			func(stack m.Containers) string { return stack[len(stack)-1:] },
			stacks,
		), "")
		return top_containers, nil
	}
}
