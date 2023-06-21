package solver

import (
	m "aoc/d04/models"
	ts "aoc/testers"
	"testing"
)

func TestD04_Containments(t *testing.T) {
	type TestCase struct {
		assignment m.SectionAssignmentPair

		is_full_overlap, is_partial_overlap bool
	}

	ts.TestThat([]TestCase{
		{m.MakeSAP(1, 5, 10, 15), false, false},
		{m.MakeSAP(1, 12, 10, 15), false, true},
		{m.MakeSAP(20, 25, 10, 15), false, false},
		{m.MakeSAP(19, 21, 10, 19), false, true},
		{m.MakeSAP(10, 12, 10, 15), true, true},
		{m.MakeSAP(50, 55, 51, 53), true, true},
		{m.MakeSAP(12, 12, 12, 12), true, true},
		{m.MakeSAP(100, 107, 107, 107), true, true},
	}, func(tc TestCase) {
		ts.AssertEqual(t, IsFullOverlap(tc.assignment), tc.is_full_overlap)
		ts.AssertEqual(t, IsPartialOverlap(tc.assignment), tc.is_partial_overlap)
	})
}
