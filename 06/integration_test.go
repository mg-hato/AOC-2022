package main

import (
	"aoc/day06/reader"
	"aoc/day06/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD06_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[string, int](t).
		ProvideReader(reading.ReadWith(reader.DatastreamBufferReader)).
		ProvideSolver(solver.FindPositionOfTheFirstMarker(4)).
		ProvideSolver(solver.FindPositionOfTheFirstMarker(14)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(7), ts.ExpectResult(19)).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(5), ts.ExpectResult(23)).
		AddTestCase("./tests/input-3.txt", ts.ExpectResult(6), ts.ExpectResult(23)).
		AddTestCase("./tests/input-4.txt", ts.ExpectResult(10), ts.ExpectResult(29)).
		AddTestCase("./tests/input-5.txt", ts.ExpectResult(11), ts.ExpectResult(26)).
		RunIntegrationTests()
}
