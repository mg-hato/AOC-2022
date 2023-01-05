package main

import (
	"aoc/argshandle"
	r "aoc/day11/reader"
	s "aoc/day11/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(r.MonkeyGraphReader),
		s.CalculateMonkeyBusiness(20, s.DivBy3),
		s.CalculateMonkeyBusiness(10_000, s.NoAdjustment),
	)
}
