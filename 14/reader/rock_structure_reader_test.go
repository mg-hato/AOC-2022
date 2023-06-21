package reader

import (
	c "aoc/common"
	m "aoc/d14/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD14_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(RockStructureReader)).
		ProvideEqualityFunction(m.RockStructureEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.RockStructureEnvelope([]m.RockStructure{
			{c.MakePair(400, 5), c.MakePair(600, 5), c.MakePair(600, 9)},
			{c.MakePair(100, 30), c.MakePair(200, 30), c.MakePair(200, 25), c.MakePair(400, 25)},
			{c.MakePair(499, 2), c.MakePair(300, 2)},
		}))).
		AddTestCase(
			"./tests/bad-1.txt",
			ts.ExpectError[m.SolverInput]("final validation", "rock formation #2", "rock line #1", "horizontal", "vertical"),
		).
		AddTestCase(
			"./tests/bad-2.txt",
			ts.ExpectError[m.SolverInput]("final validation", "rock formation #3", "overlaps", "sand source"),
		).
		AddTestCase(
			"./tests/bad-3.txt",
			ts.ExpectError[m.SolverInput]("line #4", "interpret", `"7, 7 => 9, 7"`),
		).
		RunReaderTests()

}
