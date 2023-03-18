package main

import . "aoc/functional"

// Returns the number of assignment-pairs that satisfy the passed predicate
func CountAssignmentPairsThatSatisfy(predicate func(AssignmentPair) bool, pairs []AssignmentPair) (int, error) {
	return len(Filter(predicate, pairs)), nil
}

// Return true iff
//
// - first section-range fully contains the second
//
// - OR second section-range fully contains the first
func OneFullyContainsTheOther(p AssignmentPair) bool {
	return p.first.FullyContains(p.second) || p.second.FullyContains(p.first)
}

// Return true iff the two section-ranges overlap
func SectionRangesOverlap(p AssignmentPair) bool {
	return p.first.Overlaps(p.second)
}
