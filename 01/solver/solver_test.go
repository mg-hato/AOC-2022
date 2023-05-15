package solver

import (
	m "aoc/day01/models"
	e "aoc/envelope"
	ts "aoc/testers"
	"testing"
)

func TestD01_Solver(t *testing.T) {
	type Data = e.Envelope[m.CaloryList]
	// helper functions
	cal := func(calories ...int) []int { return calories }
	cals := func(calories ...[]int) Data { return m.CreateCaloryListEnvelope(calories) }

	ts.SolverTesterForComparableResults[Data, int](t).
		ProvideSolver(CalorySumOfTop(1)).
		ProvideSolver(CalorySumOfTop(3)).
		AddTestCase(
			cals(cal(10, 20, 30)),
			ts.ExpectResult(60),
			ts.ExpectError[int]("Error", "length 1"),
		).
		AddTestCase(
			cals(cal(10), cal(5), cal(17), cal(40), cal(1, 10)),
			ts.ExpectResult(40),
			ts.ExpectResult(68),
		).
		AddTestCase(
			// sums: 10, 9, 16, 11, 11
			cals(cal(1, 2, 3, 4), cal(4, 5), cal(5, 6, 5), cal(1, 2, 2, 5, 1), cal(1, 10)),
			ts.ExpectResult(16),
			ts.ExpectResult(38),
		).
		AddTestCase(
			cals(),
			ts.ExpectError[int]("Error", "length 0"),
			ts.ExpectError[int]("Error", "length 0"),
		)
}
