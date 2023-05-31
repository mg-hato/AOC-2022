package main

import (
	"aoc/argshandle"
	"aoc/day09/reader"
	"aoc/day09/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.MotionSeriesReader),
		solver.CountPositionsVisitedByLastKnot(2),
		solver.CountPositionsVisitedByLastKnot(10),
	)
}
