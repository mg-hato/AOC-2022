package solver

import (
	m "aoc/day05/models"
	"aoc/envelope"
	f "aoc/functional"
	"strings"
)

func FollowMovingPlanWith(crate_mover CrateMover) func(envelope.Envelope[m.MovingPlan]) (string, error) {
	return func(env envelope.Envelope[m.MovingPlan]) (string, error) {
		moving_plan := env.Get()
		stacks := moving_plan.StartingContainers

		// Execute all the move instructions in order
		for i, move := range moving_plan.Moves {
			if !crate_mover.ExecuteMove(stacks, move) {
				return "", not_enough_containers_error(i+1, move)
			}
		}

		// Once finished, ensure that each stack is non-empty
		if f.Any(func(stack m.Containers) bool { return len(stack) == 0 }, stacks) {
			return "", empty_stacks_error(stacks)
		}

		// All ok, return result
		top_containers := strings.Join(f.Map(
			func(stack m.Containers) string { return stack[len(stack)-1:] },
			stacks,
		), "")
		return top_containers, nil
	}
}
