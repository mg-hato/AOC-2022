package main

import (
	"aoc/argshandle"
	"aoc/d04/reader"
	s "aoc/d04/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.SectionAssignmentsListReader),
		s.CountAssignmentPairs(s.IsFullOverlap),
		s.CountAssignmentPairs(s.IsPartialOverlap),
	)
}
