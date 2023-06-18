package reader

import (
	m "aoc/d14/models"
	e "aoc/envelope"
	f "aoc/functional"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD14_ReaderTest(t *testing.T) {
	type Data = e.Envelope[[]m.RockStructure]
	ts.ReaderTester(t, reading.ReadWith(RockStructureReader)).
		ProvideEqualityFunction(m.RockStructureEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.RockStructureEnvelope([]m.RockStructure{
			{f.MakePair(400, 5), f.MakePair(600, 5), f.MakePair(600, 9)},
			{f.MakePair(100, 30), f.MakePair(200, 30), f.MakePair(200, 25), f.MakePair(400, 25)},
			{f.MakePair(499, 2), f.MakePair(300, 2)},
		}))).
		AddTestCase(
			"./tests/bad-1.txt",
			ts.ExpectError[Data]("final validation", "rock formation #2", "rock line #1", "horizontal", "vertical"),
		).
		AddTestCase(
			"./tests/bad-2.txt",
			ts.ExpectError[Data]("final validation", "rock formation #3", "overlaps", "sand source"),
		).
		AddTestCase(
			"./tests/bad-3.txt",
			ts.ExpectError[Data]("line #4", "interpret", `"7, 7 => 9, 7"`),
		).
		RunReaderTests()

}
