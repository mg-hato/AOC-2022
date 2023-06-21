package main

import (
	"aoc/argshandle"
	"aoc/d14/reader"
	"aoc/d14/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.RockStructureReader),
		solver.CountSandUnitsUntilStop(solver.DefaultCaveSystemWithAbyss),
		solver.CountSandUnitsUntilStop(solver.DefaultCaveSystemWithFloor),
	)
}
