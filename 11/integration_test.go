package main

import (
	m "aoc/day11/models"
	r "aoc/day11/reader"
	s "aoc/day11/solver"
	e "aoc/envelope"
	"aoc/reading"
	"aoc/testers"
	"testing"
)

func TestD11_IntegrationTest(t *testing.T) {
	testers.IntegrationTesterForComparableResults[e.Envelope[[]m.Monkey], int64](t).
		ProvideReader(reading.ReadWith(r.MonkeyGraphReader)).
		ProvideSolver(s.CalculateMonkeyBusiness(20, s.DivBy3)).
		ProvideSolver(s.CalculateMonkeyBusiness(10_000, s.NoAdjustment)).
		AddTestCase(
			"./test/example.txt",
			testers.ExpectResult[int64](10_605),
			testers.ExpectResult[int64](2713310158),
		).
		AddTestCase(
			"./test/input-1.txt", // Corresponds to monkey-graph-1.jpg
			testers.ExpectResult[int64](6*5*20*20),
			testers.ExpectResult[int64](6*5*10_000*10_000),
		).
		AddTestCase(
			"./test/input-2.txt", // Corresponds to monkey-graph-2.jpg
			testers.ExpectResult[int64](101*96*20*20),
			testers.ExpectResult[int64](101*96*10_000*10_000),
		).
		RunIntegrationTests()
}
