package main

import (
	"aoc/argshandle"
	"aoc/day04/reader"
	s "aoc/day04/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.SectionAssignmentsListReader),
		s.CountAssignmentPairs(s.IsFullOverlap),
		s.CountAssignmentPairs(s.IsPartialOverlap),
	)
}
