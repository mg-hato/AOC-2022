package main

import (
	m "aoc/day02/models"
	"aoc/day02/reader"
	"aoc/day02/solver"
	e "aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD02_IntegrationTest(t *testing.T) {
	type Data = e.Envelope[[]m.Round]
	ts.IntegrationTesterForComparableResults[Data, int](t).
		ProvideReader(reading.ReadWith(reader.StategyGuideReader)).
		ProvideSolver(solver.CalculateScore(solver.ShapeBasedRoundInterpreter())).
		ProvideSolver(solver.CalculateScore(solver.OutcomeBasedRoundInterpreter())).
		AddTestCase("./tests/example.txt", ts.ExpectResult(15), ts.ExpectResult(12)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(46), ts.ExpectResult(53)).
		RunIntegrationTests()
}
