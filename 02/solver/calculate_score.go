package solver

import (
	c "aoc/common"
	m "aoc/d02/models"
)

func CalculateScore(ri RoundInterpreter) func(m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {
		return c.Sum(c.Map(ri.GetScore, input.Get())), nil
	}
}
