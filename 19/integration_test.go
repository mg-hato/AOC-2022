package main

import (
	"aoc/d19/models"
	"aoc/d19/reader"
	"aoc/d19/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD19_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[models.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.BlueprintsReader)).
		ProvideSolver(solver.AnalyseBlueprintQualities(24)).
		ProvideSolver(solver.AnalyseBlueprintQualityOfFirstN(3, 32)).
		AddTestCase(
			"./tests/example.txt",
			ts.ExpectResult(33),
			ts.ExpectResult(56*62),
		).
		RunIntegrationTests()
}
