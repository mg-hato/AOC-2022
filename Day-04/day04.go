package main

import (
	"aoc/argshandle"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		ReadListOfAssignmentPairs,
		func(pairs []AssignmentPair) int {
			return CountAssignmentPairsThatSatisfy(OneFullyContainsTheOther, pairs)
		},
		func(pairs []AssignmentPair) int {
			return CountAssignmentPairsThatSatisfy(SectionRangesOverlap, pairs)
		},
	)
}
