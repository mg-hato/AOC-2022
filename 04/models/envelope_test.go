package models

import (
	"aoc/functional"
	"aoc/testers"
	"testing"
)

func TestD04_EnvelopeTest(t *testing.T) {
	envelope := CreateSectionAssignmentPairsEnvelope([]SectionAssignmentPair{MakeSAP(1, 5, 10, 15), MakeSAP(7, 7, 7, 10)})
	pairs := envelope.Get()
	pairs[0] = MakeSAP(0, 0, 0, 0)
	pairs[1].Second.Right = 100
	testers.AssertEqualWithEqFunc(
		t,
		envelope.Get(),
		[]SectionAssignmentPair{MakeSAP(1, 5, 10, 15), MakeSAP(7, 7, 7, 10)},
		functional.ArrayEqual[SectionAssignmentPair],
	)
}
