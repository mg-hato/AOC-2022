package main

import (
	"aoc/argshandle"
)

func main() {

	argshandle.AoC2022DefaultProgram(
		ReadListOfContents,
		func(lc ListOfContents) int { return SumOfPriorities(lc, FindRepeatedItems) },
		func(lc ListOfContents) int { return SumOfPriorities(lc, FindGroupBadges) },
	)
}
