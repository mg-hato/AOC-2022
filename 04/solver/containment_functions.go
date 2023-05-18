package solver

import (
	m "aoc/day04/models"
	"aoc/functional"
)

// Returns true if and only if given coverage covers all of the passed sections
func contains(coverage m.Coverage, sections ...int) bool {
	return functional.All(
		func(section int) bool { return coverage.Left <= section && section <= coverage.Right },
		sections,
	)
}

// Returns true if and only if one elf covers entirely the coverage of the other elf
func IsFullOverlap(assignment m.SectionAssignmentPair) bool {
	return contains(assignment.First, assignment.Second.Left, assignment.Second.Right) ||
		contains(assignment.Second, assignment.First.Left, assignment.First.Right)
}

// Returns true if and only if there exists a section that is covered by both elves
func IsPartialOverlap(assignment m.SectionAssignmentPair) bool {
	return contains(assignment.First, assignment.Second.Left) || contains(assignment.First, assignment.Second.Right) ||
		contains(assignment.Second, assignment.First.Left) || contains(assignment.Second, assignment.First.Right)
}
