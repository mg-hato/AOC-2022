package reader

import (
	m "aoc/day01/models"
	e "aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD01_ReaderTest(t *testing.T) {
	type Data = e.Envelope[m.CaloryList]
	cal := func(calories ...int) []int { return calories }
	cals := func(calories ...[]int) Data { return m.CreateCaloryListEnvelope(calories) }

	ts.ReaderTester(t, reading.ReadWith(CaloryListReader)).
		ProvideEqualityFunction(m.CaloryListEqualityFunc).
		AddTestCase(
			"./test/bad-input-1.txt",
			ts.ExpectError[Data]("empty"),
		).
		AddTestCase(
			"./test/bad-input-2.txt",
			ts.ExpectError[Data]("line #5", "012"),
		).
		AddTestCase(
			"./test/bad-input-3.txt",
			ts.ExpectError[Data]("line #4", "1 2 3"),
		).
		AddTestCase(
			"./test/bad-input-4.txt",
			ts.ExpectError[Data]("line #7", "19A"),
		).
		AddTestCase(
			"./test/input-1.txt",
			ts.ExpectResult(cals(cal(10), cal(20))),
		).
		AddTestCase(
			"./test/input-2.txt",
			ts.ExpectResult(cals(cal(10, 13, 17))),
		).
		AddTestCase(
			"./test/input-3.txt",
			ts.ExpectResult(cals(cal(1))),
		).
		AddTestCase(
			"./test/input-4.txt",
			ts.ExpectResult(cals(cal(0))),
		).
		RunReaderTests()
}
