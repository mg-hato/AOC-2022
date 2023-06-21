package solver

import (
	c "aoc/common"
	m "aoc/d04/models"
)

func CountAssignmentPairs(predicate func(m.SectionAssignmentPair) bool) func(m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {
		return len(c.Filter(predicate, input.Get())), nil
	}
}
