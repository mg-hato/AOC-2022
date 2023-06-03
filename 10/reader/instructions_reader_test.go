package reader

import (
	m "aoc/day10/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD10_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(InstructionsReader)).
		ProvideEqualityFunction(m.InstructionsEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.InstructionsEnvelope([]m.Instruction{
			m.Addx(10),
			m.Noop(),
			m.Noop(),
			m.Addx(-17),
			m.Noop(),
			m.Addx(0),
		}))).
		RunReaderTests()
}
