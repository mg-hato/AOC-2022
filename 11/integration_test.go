package main

import (
	m "aoc/d11/models"
	r "aoc/d11/reader"
	s "aoc/d11/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD11_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int64](t).
		ProvideReader(reading.ReadWith(r.MonkeyGraphReader)).
		ProvideSolver(s.CalculateMonkeyBusiness(20, s.DivBy3)).
		ProvideSolver(s.CalculateMonkeyBusiness(10_000, s.NoAdjustment)).
		AddTestCase(
			"./test/example.txt",
			ts.ExpectResult[int64](10_605),
			ts.ExpectResult[int64](2713310158),
		).
		AddTestCase(
			"./test/input-1.txt", // Corresponds to monkey-graph-1.jpg
			ts.ExpectResult[int64](6*5*20*20),
			ts.ExpectResult[int64](6*5*10_000*10_000),
		).
		AddTestCase(
			"./test/input-2.txt", // Corresponds to monkey-graph-2.jpg
			ts.ExpectResult[int64](101*96*20*20),
			ts.ExpectResult[int64](101*96*10_000*10_000),
		).
		RunIntegrationTests()
}
