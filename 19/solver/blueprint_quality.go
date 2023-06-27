package solver

import (
	c "aoc/common"
	m "aoc/d19/models"
)

func blueprint_quality(time int) func(m.Blueprint) int {
	return func(bp m.Blueprint) int {
		pq := c.PriorityQueue(func(lhs, rhs m.State) bool {
			return lhs.Estimated > rhs.Estimated
		})
		pq.Add(bp.MakeInitialState(time))
		next_state_func := bp.GetNextStateFunction()
		achieved := 0
		for !pq.IsEmpty() {
			state, _ := pq.Pop()
			if achieved < state.Estimated {
				achieved = c.Max(achieved, state.GetGeodeCrackedUntilTimeout())
				for _, next_state := range next_state_func(state) {
					if achieved < next_state.Estimated {
						pq.Add(next_state)
					}
				}
			}
		}
		return achieved
	}
}
