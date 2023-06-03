package main

import (
	"aoc/argshandle"
	"aoc/day10/reader"
	"aoc/day10/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.InstructionsReader),
		solver.SimulateProgram(solver.SignalStrengthAnalyser(20, 60, 100, 140, 180, 220)),
		solver.SimulateProgram(solver.ImageDrawerAnalyser(40, 6)),
	)
}
