package main

import (
	"aoc/d05/models"
	"aoc/d05/reader"
	"aoc/d05/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD05_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[models.SolverInput, string](t).
		ProvideReader(reading.ReadWith(reader.MovingPlanReader)).
		ProvideSolver(solver.FollowMovingPlanWith(solver.CrateMover9000())).
		ProvideSolver(solver.FollowMovingPlanWith(solver.CrateMover9001())).
		AddTestCase("./tests/example.txt", ts.ExpectResult("CMZ"), ts.ExpectResult("MCD")).
		RunIntegrationTests()
}
