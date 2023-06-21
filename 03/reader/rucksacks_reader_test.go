package reader

import (
	m "aoc/d03/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD03_ReaderTest(t *testing.T) {
	data := func(rucksacks ...m.Rucksack) m.SolverInput { return m.CreateRucksacksEnvelope(rucksacks) }

	ts.ReaderTester(t, reading.ReadWith(RucksacksReader)).
		ProvideEqualityFunction(m.RucksacksEnvelopeEqualityFunction).
		AddTestCase("./tests/bad-input-1.txt", ts.ExpectError[m.SolverInput]("number of", "split", "groups of three")).
		AddTestCase("./tests/bad-input-2.txt", ts.ExpectError[m.SolverInput]("line #3", "rucksack must contain", "even number", "JJJIKLM")).
		AddTestCase("./tests/bad-input-3.txt", ts.ExpectError[m.SolverInput]("line #2", "interpret", "cdef1b")).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(data("abcdef", "ghij", "klmnopqrst"))).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(data("AB", "xyzi", "pppq", "wyWY", "iiiijj", "ag"))).
		RunReaderTests()
}
