package main

import (
	"aoc/argshandle"
	"aoc/d10/reader"
	"aoc/d10/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.InstructionsReader),
		solver.SimulateProgram(solver.SignalStrengthAnalyser(20, 60, 100, 140, 180, 220)),
		solver.SimulateProgram(solver.ImageDrawerAnalyser(40, 6)),
	)
}
