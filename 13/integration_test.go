package main

import (
	"aoc/d13/models"
	"aoc/d13/reader"
	"aoc/d13/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD13_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[models.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.PacketReader)).
		ProvideSolver(solver.CountOrderedPacketPairs).
		ProvideSolver(solver.ExtractDecoderKey).
		AddTestCase("./tests/example.txt", ts.ExpectResult(13), ts.ExpectResult(140)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(14), ts.ExpectResult(96)).
		RunIntegrationTests()
}
