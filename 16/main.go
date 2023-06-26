package main

import (
	"aoc/argshandle"
	"aoc/d16/reader"
	"aoc/d16/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.ValvesReader),
		solver.FindMaxPressureRelease(1, 30),
		solver.FindMaxPressureRelease(2, 26),
	)
}
