package main

import (
	m "aoc/d02/models"
	"aoc/d02/reader"
	"aoc/d02/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD02_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.StategyGuideReader)).
		ProvideSolver(solver.CalculateScore(solver.ShapeBasedRoundInterpreter())).
		ProvideSolver(solver.CalculateScore(solver.OutcomeBasedRoundInterpreter())).
		AddTestCase("./tests/example.txt", ts.ExpectResult(15), ts.ExpectResult(12)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(46), ts.ExpectResult(53)).
		RunIntegrationTests()
}
