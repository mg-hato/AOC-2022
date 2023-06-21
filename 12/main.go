package main

import (
	"aoc/argshandle"
	"aoc/d12/reader"
	"aoc/d12/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.TerrainReader),
		solver.CalculateDistance(solver.StartingPositionDistancePicker('S')),
		solver.CalculateDistance(solver.StartingPositionDistancePicker('S', 'a')),
	)
}
