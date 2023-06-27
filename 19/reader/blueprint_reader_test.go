package reader

import (
	m "aoc/d19/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD19_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(BlueprintsReader)).
		ProvideEqualityFunction(m.BlueprintsEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.BlueprintsEnvelope(
			m.MakeBlueprint(1, 1, 5, 1, 4, 40, 20),
			m.MakeBlueprint(2, 2, 3, 5, 11, 20, 40),
			m.MakeBlueprint(3, 3, 2, 3, 9, 9, 18),
		))).
		AddTestCase("./tests/invalid_ids.txt", ts.ExpectError[m.SolverInput](
			"validation",
			"[1, 3]",
			"ascending",
		)).
		RunReaderTests()

}
