package main

import (
	"aoc/d14/models"
	"aoc/d14/reader"
	"aoc/d14/solver"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD14_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[envelope.Envelope[[]models.RockStructure], int](t).
		ProvideReader(reading.ReadWith(reader.RockStructureReader)).
		ProvideSolver(solver.CountSandUnitsUntilStop(solver.DefaultCaveSystemWithAbyss)).
		ProvideSolver(solver.CountSandUnitsUntilStop(solver.DefaultCaveSystemWithFloor)).
		AddTestCase("./tests/example.txt", ts.ExpectResult(24), ts.ExpectResult(93)).
		RunIntegrationTests()
}
