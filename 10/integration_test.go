package main

import (
	m "aoc/d10/models"
	"aoc/d10/reader"
	"aoc/d10/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD10_IntegrationTest(t *testing.T) {
	type Result = m.AnalyserReport
	ts.IntegrationTester[m.SolverInput, Result](t).
		ProvideReader(reading.ReadWith(reader.InstructionsReader)).
		ProvideSolver(solver.SimulateProgram(solver.SignalStrengthAnalyser(20, 60, 100, 140, 180, 220))).
		ProvideSolver(solver.SimulateProgram(solver.ImageDrawerAnalyser(40, 6))).
		ProvideEqualityFunctionForResults(m.AnalyserReportEqualityFunction).
		AddTestCase(
			"./tests/example.txt",
			ts.ExpectResult(m.SignalStrengthReport(13_140)),
			ts.ExpectResult(m.ImageReport([][]rune{
				[]rune("##..##..##..##..##..##..##..##..##..##.."),
				[]rune("###...###...###...###...###...###...###."),
				[]rune("####....####....####....####....####...."),
				[]rune("#####.....#####.....#####.....#####....."),
				[]rune("######......######......######......####"),
				[]rune("#######.......#######.......#######....."),
			})),
		).
		RunIntegrationTests()
}
