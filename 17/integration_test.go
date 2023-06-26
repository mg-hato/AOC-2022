package main

import (
	"aoc/d17/reader"
	"aoc/d17/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD17_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[string, int64](t).
		ProvideReader(reading.ReadWith(reader.JetPatternReader)).
		ProvideSolver(solver.GetHeightAfterNRocks(2022)).
		ProvideSolver(solver.GetHeightAfterNRocks(1_000_000_000_000)).
		AddTestCase(
			"./tests/example.txt",
			ts.ExpectResult[int64](3068),
			ts.ExpectResult[int64](1_514_285_714_288),
		).
		RunIntegrationTests()
}
