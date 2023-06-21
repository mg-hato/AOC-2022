package solver

import (
	c "aoc/common"
	m "aoc/d09/models"
)

func CountPositionsVisitedByLastKnot(number_of_knots int) func(m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {
		var knots []m.Position = c.Repeat(c.MakePair(0, 0), number_of_knots)

		visited := map[m.Position]bool{c.MakePair(0, 0): true}
		for _, motion := range input.Get() {
			for i := 0; i < motion.Steps; i++ {
				move_knots(motion.Direction, knots)
				visited[knots[number_of_knots-1]] = true
			}
		}

		return c.Count(c.GetValues(visited), c.Identity[bool]), nil
	}
}
