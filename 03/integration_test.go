package main

import (
	"aoc/day03/models"
	"aoc/day03/reader"
	"aoc/day03/solver"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD03_IntegrationTest(t *testing.T) {
	type Data = envelope.Envelope[[]models.Rucksack]
	ts.IntegrationTesterForComparableResults[Data, int](t).
		ProvideReader(reading.ReadWith(reader.RucksacksReader)).
		ProvideSolver(solver.SumItemPriorities(solver.CompartmentDuplicateItemLocator())).
		ProvideSolver(solver.SumItemPriorities(solver.BadgeItemLocator())).
		AddTestCase("./tests/example.txt", ts.ExpectResult(157), ts.ExpectResult(70)).
		RunIntegrationTests()
}
