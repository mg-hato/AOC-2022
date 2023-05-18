package models

import (
	e "aoc/envelope"
	f "aoc/functional"
	"fmt"
	"strings"
)

type SectionAssignmentPairsEnvelope struct {
	assignment_pairs []SectionAssignmentPair
}

func CreateSectionAssignmentPairsEnvelope(pairs []SectionAssignmentPair) e.Envelope[[]SectionAssignmentPair] {
	return &SectionAssignmentPairsEnvelope{pairs}
}

func (env SectionAssignmentPairsEnvelope) Get() []SectionAssignmentPair {
	data := make([]SectionAssignmentPair, len(env.assignment_pairs))
	copy(data, env.assignment_pairs)
	return data
}

func (env SectionAssignmentPairsEnvelope) String() string {
	return fmt.Sprintf(
		"SAP-Envelope[%s]",
		strings.Join(
			f.Map(func(sap SectionAssignmentPair) string { return fmt.Sprint(sap) }, env.assignment_pairs),
			", ",
		))
}

func SectionAssignmentPairsEnvelopeEqualityFunction(lhs, rhs e.Envelope[[]SectionAssignmentPair]) bool {
	return f.ArrayEqual(lhs.Get(), rhs.Get())
}
