package main

import (
	"aoc/argshandle"
	r "aoc/d03/reader"
	s "aoc/d03/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(r.RucksacksReader),
		s.SumItemPriorities(s.CompartmentDuplicateItemLocator()),
		s.SumItemPriorities(s.BadgeItemLocator()),
	)
}
