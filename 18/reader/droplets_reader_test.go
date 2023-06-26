package reader

import (
	m "aoc/d18/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD18_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(DropletsReader)).
		ProvideEqualityFunction(m.DropletsEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.DropletsEnvelope(
			m.MakeDroplet(10, 10, 10),
			m.MakeDroplet(2, 5, 7),
			m.MakeDroplet(10, 11, 10),
			m.MakeDroplet(3, 6, 8),
		))).
		RunReaderTests()
}
