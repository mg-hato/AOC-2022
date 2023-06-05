package reader

import (
	m "aoc/d12/models"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD12_ReaderTest(t *testing.T) {
	type Data = envelope.Envelope[m.Terrain]
	ts.ReaderTester(t, reading.ReadWith(TerrainReader)).
		ProvideEqualityFunction(m.TerrainEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.TerrainEnvelope("SabcdefghijklmnopqrstuvwxyzE"))).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(m.TerrainEnvelope("xyz", "abc", "EiS"))).
		AddTestCase("./tests/bad-input-1.txt",
			ts.ExpectError[Data]("different lengths", "across", "rows"),
		).
		AddTestCase("./tests/bad-input-2.txt",
			ts.ExpectError[Data]("exactly one", "finish position", "there are 2 instead"),
		).
		AddTestCase("./tests/bad-input-3.txt",
			ts.ExpectError[Data]("exactly one", "start", "position", "there are 2 instead"),
		).
		AddTestCase("./tests/bad-input-4.txt",
			ts.ExpectError[Data]("exactly one", "start", "position", "there are 0 instead"),
		).
		AddTestCase("./tests/bad-input-5.txt",
			ts.ExpectError[Data]("line #3", "could", "not", "interpret"),
		).
		RunReaderTests()
}
