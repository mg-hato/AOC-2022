package main

import (
	"aoc/d09/models"
	"aoc/d09/reader"
	"aoc/d09/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD09_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[models.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.MotionSeriesReader)).
		ProvideSolver(solver.CountPositionsVisitedByLastKnot(2)).
		ProvideSolver(solver.CountPositionsVisitedByLastKnot(10)).
		AddTestCase("./tests/example-1.txt", ts.ExpectResult(13), ts.ExpectResult(1)).
		AddTestCase("./tests/example-2.txt", ts.ExpectResult(88), ts.ExpectResult(36)).
		RunIntegrationTests()

	// For "example-2.txt" see "example-2-explained.txt" explaining the result of 88
}
