package reader

import (
	m "aoc/day08/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD08_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(ForestReader)).
		ProvideEqualityFunction(m.ForestEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.ForestEnvelope(m.Forest{
			{1, 2, 3, 4},
			{9, 9, 8, 9},
			{7, 7, 7, 7},
			{1, 2, 2, 1},
		}))).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(m.ForestEnvelope(m.Forest{
			{1, 2, 3, 4, 4, 5, 6, 7, 8, 9},
			{9, 9, 9, 8, 7, 7, 6, 1, 2, 9},
			{1, 2, 5, 5, 5, 3, 4, 5, 6, 7},
			{9, 9, 9, 2, 3, 4, 2, 2, 2, 7},
			{1, 1, 1, 3, 4, 8, 9, 9, 9, 4},
			{4, 4, 4, 9, 9, 7, 7, 1, 1, 2},
			{1, 1, 2, 1, 1, 1, 2, 1, 1, 3},
		}))).
		RunReaderTests()
}
