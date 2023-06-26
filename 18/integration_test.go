package main

import (
	m "aoc/d18/models"
	"aoc/d18/reader"
	"aoc/d18/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD18_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.DropletsReader)).
		ProvideSolver(solver.CountAreaOfDropletSurfaces).
		ProvideSolver(solver.CountAreaOfWaterAdjacentSurfaces).
		AddTestCase("./tests/example.txt", ts.ExpectResult(64), ts.ExpectResult(58)).
		RunIntegrationTests()

}
