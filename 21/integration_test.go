package main

import (
	"aoc/d21/models"
	"aoc/d21/reader"
	"aoc/d21/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD21_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[models.SolverInput, int64](t).
		ProvideReader(reading.ReadWith(reader.MonkeyJobsReader)).
		ProvideSolver(solver.SolveRoot).
		ProvideSolver(solver.SolveHumn).
		AddTestCase("./tests/example.txt", ts.ExpectResult[int64](152), ts.ExpectResult[int64](301)).
		RunIntegrationTests()
}
