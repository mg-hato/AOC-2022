package main

import (
	"aoc/argshandle"
	r "aoc/d01/reader"
	s "aoc/d01/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(r.CaloryListReader),
		s.CalorySumOfTop(1),
		s.CalorySumOfTop(3),
	)
}
