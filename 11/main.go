package main

import (
	"aoc/argshandle"
	r "aoc/d11/reader"
	s "aoc/d11/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(r.MonkeyGraphReader),
		s.CalculateMonkeyBusiness(20, s.DivBy3),
		s.CalculateMonkeyBusiness(10_000, s.NoAdjustment),
	)
}
