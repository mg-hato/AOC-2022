package main

import (
	"aoc/d23/models"
	"aoc/d23/reader"
	"aoc/d23/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD22_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[models.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.GroveMapReader)).
		ProvideSolver(solver.CountFreeSpacesInEncapsulatingRegion(10)).
		ProvideSolver(solver.FirstRoundWhenNoElfMoves(100)).
		AddTestCase("./tests/example.txt", ts.ExpectResult(110), ts.ExpectResult(20)).
		RunIntegrationTests()
}
