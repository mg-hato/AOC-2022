package main

import (
	"aoc/day08/models"
	"aoc/day08/reader"
	"aoc/day08/solver"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD08_IntegrationTest(t *testing.T) {
	type Data = envelope.Envelope[models.Forest]
	ts.IntegrationTesterForComparableResults[Data, int](t).
		ProvideReader(reading.ReadWith(reader.ForestReader)).
		ProvideSolver(solver.AnalyseForest(solver.VisibilityTreeLineAnalyser)).
		ProvideSolver(solver.AnalyseForest(solver.ScenicScoreTreeLineAnalyser)).
		AddTestCase("./tests/example.txt", ts.ExpectResult(21), ts.ExpectResult(8)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(62), ts.ExpectResult(0)).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(43), ts.ExpectResult(0)).
		AddTestCase("./tests/input-3.txt", ts.ExpectResult(65), ts.ExpectResult(420)).
		RunIntegrationTests()
}
