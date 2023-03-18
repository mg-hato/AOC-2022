package main

import (
	"aoc/argshandle"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		ReadListOfAssignmentPairs,
		func(pairs []AssignmentPair) (int, error) {
			return CountAssignmentPairsThatSatisfy(OneFullyContainsTheOther, pairs)
		},
		func(pairs []AssignmentPair) (int, error) {
			return CountAssignmentPairsThatSatisfy(SectionRangesOverlap, pairs)
		},
	)
}
