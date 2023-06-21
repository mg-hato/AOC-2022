package main

import (
	m "aoc/d03/models"
	"aoc/d03/reader"
	"aoc/d03/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD03_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.RucksacksReader)).
		ProvideSolver(solver.SumItemPriorities(solver.CompartmentDuplicateItemLocator())).
		ProvideSolver(solver.SumItemPriorities(solver.BadgeItemLocator())).
		AddTestCase("./tests/example.txt", ts.ExpectResult(157), ts.ExpectResult(70)).
		RunIntegrationTests()
}
