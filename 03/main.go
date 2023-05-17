package main

import (
	"aoc/argshandle"
	r "aoc/day03/reader"
	s "aoc/day03/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(r.RucksacksReader),
		s.SumItemPriorities(s.CompartmentDuplicateItemLocator()),
		s.SumItemPriorities(s.BadgeItemLocator()),
	)
}
