package main

import (
	m "aoc/d08/models"
	"aoc/d08/reader"
	"aoc/d08/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD08_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.ForestReader)).
		ProvideSolver(solver.AnalyseForest(solver.VisibilityTreeLineAnalyser)).
		ProvideSolver(solver.AnalyseForest(solver.ScenicScoreTreeLineAnalyser)).
		AddTestCase("./tests/example.txt", ts.ExpectResult(21), ts.ExpectResult(8)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(62), ts.ExpectResult(0)).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(43), ts.ExpectResult(0)).
		AddTestCase("./tests/input-3.txt", ts.ExpectResult(65), ts.ExpectResult(420)).
		RunIntegrationTests()
}
