package models

import (
	c "aoc/common"
	"fmt"
	"strings"
)

type SectionAssignmentPairsEnvelope struct {
	assignment_pairs []SectionAssignmentPair
}

func CreateSectionAssignmentPairsEnvelope(pairs []SectionAssignmentPair) c.Envelope[[]SectionAssignmentPair] {
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
			c.Map(func(sap SectionAssignmentPair) string { return fmt.Sprint(sap) }, env.assignment_pairs),
			", ",
		))
}

func SectionAssignmentPairsEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]SectionAssignmentPair]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
