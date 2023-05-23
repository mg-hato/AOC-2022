package main

import (
	"aoc/day07/models"
	"aoc/day07/reader"
	s "aoc/day07/solver"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD07_IntegrationTest(t *testing.T) {
	type Data = envelope.Envelope[[]models.Command]
	spec := s.SimpleFilesystemSpec(70_000_000, 30_000_000)
	ts.IntegrationTesterForComparableResults[Data, int64](t).
		ProvideReader(reading.ReadWith(reader.TerminalOutputReader)).
		ProvideSolver(s.AnalyseFilesystem(spec, s.SumDirectoriesOfSizeAtMost(100_000))).
		ProvideSolver(s.AnalyseFilesystem(spec, s.FindSmallestDirectoryEnablingUpdate())).
		AddTestCase(
			"./tests/example.txt",
			ts.ExpectResult[int64](95_437),
			ts.ExpectResult[int64](24_933_642),
		).
		AddTestCase(
			"./tests/input-1.txt",
			ts.ExpectResult[int64](199_002),
			// Memory occupied is 70 million + 1 units which is by 1 greater than total memory
			ts.ExpectError[int64]("free memory", "negative number", "-1"),
		).
		AddTestCase(
			"./tests/input-2.txt",
			ts.ExpectResult[int64](12_000),
			ts.ExpectError[int64]("enough", "free memory", "update"),
		).
		RunIntegrationTests()
}
