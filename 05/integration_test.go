package main

import (
	"aoc/day05/models"
	"aoc/day05/reader"
	"aoc/day05/solver"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD05_IntegrationTest(t *testing.T) {
	type Data = envelope.Envelope[models.MovingPlan]
	ts.IntegrationTesterForComparableResults[Data, string](t).
		ProvideReader(reading.ReadWith(reader.MovingPlanReader)).
		ProvideSolver(solver.FollowMovingPlanWith(solver.CrateMover9000())).
		ProvideSolver(solver.FollowMovingPlanWith(solver.CrateMover9001())).
		AddTestCase("./tests/example.txt", ts.ExpectResult("CMZ"), ts.ExpectResult("MCD")).
		RunIntegrationTests()
}
