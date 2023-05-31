package reader

import (
	m "aoc/day09/models"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD09_ReaderTest(t *testing.T) {
	up, down, left, right := m.UP, m.DOWN, m.LEFT, m.RIGHT
	ts.ReaderTester(t, reading.ReadWith(MotionSeriesReader)).
		ProvideEqualityFunction(m.MotionSeriesEnvelopeEqualityFunc).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.MotionSeriesEnvelope(m.MotionSeries{}))).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(m.MotionSeriesEnvelope(m.MotionSeries{
			m.MakeMotion(15, down),
		}))).
		AddTestCase("./tests/input-3.txt", ts.ExpectResult(m.MotionSeriesEnvelope(m.MotionSeries{
			m.MakeMotion(1, up),
			m.MakeMotion(7, down),
			m.MakeMotion(10, left),
			m.MakeMotion(16, right),
		}))).
		AddTestCase(
			"./tests/bad-input-1.txt",
			ts.ExpectError[envelope.Envelope[m.MotionSeries]](`"R 0"`, "line #4", "interpret"),
		).
		RunReaderTests()
}
