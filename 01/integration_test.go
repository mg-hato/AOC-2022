package main

import (
	m "aoc/d01/models"
	r "aoc/d01/reader"
	s "aoc/d01/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD01_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int](t).
		ProvideReader(reading.ReadWith(r.CaloryListReader)).
		ProvideSolver(s.CalorySumOfTop(1)).
		ProvideSolver(s.CalorySumOfTop(3)).
		AddTestCase(
			"./tests/example.txt",
			ts.ExpectResult(24_000),
			ts.ExpectResult(45_000),
		).
		AddTestCase(
			"./tests/input-1.txt",
			ts.ExpectResult(360),
			ts.ExpectError[int]("length 2"), // there are only two elves/bags, so sum-up of top 3 should return an error
		).
		AddTestCase(
			// sums per bag: 150, 200, 160, 130, 210, 160
			"./tests/input-2.txt",
			ts.ExpectResult(210),
			ts.ExpectResult(210+200+160),
		)
}
