package solver

import (
	c "aoc/common"
	"aoc/d19/models"
)

func AnalyseBlueprintQualityOfFirstN(n int, time int) func(models.SolverInput) (int, error) {
	return func(envelope models.SolverInput) (int, error) {
		blueprint_quality_func := blueprint_quality(time)
		return c.Foldl(
			func(L, R int) int { return L * R },
			c.Map(blueprint_quality_func, c.Take(n, envelope.Get())),
			1,
		), nil
	}
}
