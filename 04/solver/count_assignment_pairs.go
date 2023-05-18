package solver

import (
	m "aoc/day04/models"
	e "aoc/envelope"
	"aoc/functional"
)

func CountAssignmentPairs(predicate func(m.SectionAssignmentPair) bool) func(e.Envelope[[]m.SectionAssignmentPair]) (int, error) {
	return func(envelope e.Envelope[[]m.SectionAssignmentPair]) (int, error) {
		return len(functional.Filter(predicate, envelope.Get())), nil
	}
}
