package main

import "fmt"

// Representation of all the sections starting from section #`left` to #`right` (inclusive)
type SectionRange struct {
	left, right int
}

// A pair of elves with their sections they need to clean
type AssignmentPair struct {
	first, second SectionRange
}

func (ap AssignmentPair) String() string {
	return fmt.Sprintf("(%s,%s)", ap.first, ap.second)
}

func (s SectionRange) String() string {
	return fmt.Sprintf("%d-%d", s.left, s.right)
}

// Returns true iff `lhs` section-range fully contains `rhs`
func (lhs SectionRange) FullyContains(rhs SectionRange) bool {
	return lhs.Contains(rhs.left) && lhs.Contains(rhs.right)
}

// Return true iff section-ranges lhs & rhs overlap
func (lhs SectionRange) Overlaps(rhs SectionRange) bool {
	return lhs.Contains(rhs.left) || lhs.Contains(rhs.right) || rhs.FullyContains(lhs)
}

// Return true iff section `i` is contained in section-rage `s`
func (s SectionRange) Contains(i int) bool {
	return s.left <= i && i <= s.right
}
