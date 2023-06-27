package solver

import (
	c "aoc/common"
	m "aoc/d19/models"
)

func AnalyseBlueprintQualities(time int) func(m.SolverInput) (int, error) {
	return func(envelope m.SolverInput) (int, error) {
		quality_func := blueprint_quality(time)

		return c.Sum(c.Map(
			func(bp m.Blueprint) int { return bp.ID * quality_func(bp) },
			envelope.Get(),
		)), nil
	}
}
