package main

import (
	"aoc/d12/models"
	"aoc/d12/reader"
	"aoc/d12/solver"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD12_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[envelope.Envelope[models.Terrain], int](t).
		ProvideReader(reading.ReadWith(reader.TerrainReader)).
		ProvideSolver(solver.CalculateDistance(solver.StartingPositionDistancePicker('S'))).
		ProvideSolver(solver.CalculateDistance(solver.StartingPositionDistancePicker('S', 'a'))).
		AddTestCase("./tests/example.txt", ts.ExpectResult(31), ts.ExpectResult(29)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(40+1+40), ts.ExpectResult(24+1+40)).
		RunIntegrationTests()

}
