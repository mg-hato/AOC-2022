package main

import (
	"aoc/argshandle"
	"aoc/d14/reader"
	"aoc/d14/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.RockStructureReader),
		solver.CountSandUnitsUntilStop(solver.DefaultCaveSystemWithAbyss),
		solver.CountSandUnitsUntilStop(solver.DefaultCaveSystemWithFloor),
	)
}
