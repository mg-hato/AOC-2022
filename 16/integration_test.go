package main

import (
	m "aoc/d16/models"
	"aoc/d16/reader"
	"aoc/d16/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD16_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.ValvesReader)).
		ProvideSolver(solver.FindMaxPressureRelease(1, 30)).
		ProvideSolver(solver.FindMaxPressureRelease(2, 26)).
		AddTestCase("./tests/example.txt", ts.ExpectResult(1651), ts.ExpectResult(1707)).
		AddTestCase("./tests/test-example.txt", ts.ExpectResult(366), ts.ExpectResult(326)).
		RunIntegrationTests()
}
