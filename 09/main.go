package main

import (
	"aoc/argshandle"
	"aoc/d09/reader"
	"aoc/d09/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.MotionSeriesReader),
		solver.CountPositionsVisitedByLastKnot(2),
		solver.CountPositionsVisitedByLastKnot(10),
	)
}
