package solver

import (
	c "aoc/common"
	"aoc/d16/models"
)

func FindMaxPressureRelease(number_of_agents, deadline int) func(models.SolverInput) (int, error) {
	return func(input models.SolverInput) (int, error) {

		valves := input.Get()

		g := make_graph(valves, deadline)

		achieved := 0
		pq := c.PriorityQueue(func(lhs, rhs state) bool {
			return lhs.oprp > rhs.oprp
		})

		pq.Add(make_initial_state(valves, number_of_agents))

		for !pq.IsEmpty() {
			S, _ := pq.Pop()
			if S.oprp > achieved {
				achieved = c.Max(achieved, g.get_secured_pressure_release(S))
				pq.MultiAdd(g.expand_state(S)...)
			}
		}
		return achieved, nil
	}
}
