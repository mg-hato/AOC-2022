package main

import (
	"aoc/argshandle"
	r "aoc/day01/reader"
	s "aoc/day01/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(r.CaloryListReader),
		s.CalorySumOfTop(1),
		s.CalorySumOfTop(3),
	)
}
