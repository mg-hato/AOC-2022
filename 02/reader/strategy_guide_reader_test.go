package reader

import (
	m "aoc/d02/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD02_ReaderTest(t *testing.T) {

	// helper function
	r := func(left m.LeftSymbol, right m.RightSymbol) m.Round { return m.Round{Left: left, Right: right} }

	// shortcut for ABC & XYZ symbols
	x, y, z := m.X, m.Y, m.Z
	a, b, c := m.A, m.B, m.C

	ts.ReaderTester(t, reading.ReadWith(StategyGuideReader)).
		ProvideEqualityFunction(m.RoundsEnvelopeEqualityFunction).
		AddTestCase("./tests/bad-input-1.txt", ts.ExpectError[m.SolverInput]("empty")).
		AddTestCase("./tests/bad-input-2.txt", ts.ExpectError[m.SolverInput]("line #3", "B C")).
		AddTestCase("./tests/bad-input-3.txt", ts.ExpectError[m.SolverInput]("line #4", "1 0")).
		AddTestCase("./tests/input-1.txt",
			ts.ExpectResult(m.CreateRoundsEnvelope([]m.Round{
				r(a, z),
				r(b, y),
				r(c, x),
				r(c, z),
				r(b, x),
			}))).
		AddTestCase("./tests/input-2.txt",
			ts.ExpectResult(m.CreateRoundsEnvelope([]m.Round{
				r(a, x),
			}))).
		RunReaderTests()
}
