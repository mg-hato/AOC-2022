package main

import (
	"aoc/d15/models"
	"aoc/d15/reader"
	"aoc/d15/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD15_IntegrationTest(t *testing.T) {

	ts.IntegrationTesterForComparableResults[models.SolverInput, int64](t).
		ProvideReader(reading.ReadWith(reader.SensorReportsReader)).
		ProvideSolver(solver.BeaconExclusionCount(9)).
		ProvideSolver(solver.BeaconExclusionCount(10)).
		ProvideSolver(solver.BeaconExclusionCount(11)).
		ProvideSolver(solver.DistressBeaconTuningFrequencyFinder(20)).
		AddTestCase("./tests/example.txt",
			ts.ExpectResult[int64](25),
			ts.ExpectResult[int64](26),
			ts.ExpectResult[int64](28),
			ts.ExpectResult[int64](14*20+11),
		).
		RunIntegrationTests()

}
