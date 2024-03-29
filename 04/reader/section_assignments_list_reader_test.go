package reader

import (
	m "aoc/d04/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD04_ReaderTest(t *testing.T) {
	data := func(pairs ...m.SectionAssignmentPair) m.SolverInput {
		return m.CreateSectionAssignmentPairsEnvelope(pairs)
	}

	ts.ReaderTester(t, reading.ReadWith(SectionAssignmentsListReader)).
		ProvideEqualityFunction(m.SectionAssignmentPairsEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(data(
			m.MakeSAP(1, 2, 3, 4),
		))).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(data(
			m.MakeSAP(101, 102, 99, 2_000),
			m.MakeSAP(55, 55, 55, 900),
			m.MakeSAP(1, 7, 20, 20),
		))).
		AddTestCase("./tests/bad-input-1.txt", ts.ExpectError[m.SolverInput]("line #3", "invalid", "order")).
		AddTestCase("./tests/bad-input-2.txt", ts.ExpectError[m.SolverInput]("line #2", "invalid", "order")).
		RunReaderTests()
}
