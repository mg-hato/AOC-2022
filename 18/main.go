package main

import (
	"aoc/argshandle"
	"aoc/d18/reader"
	"aoc/d18/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.DropletsReader),
		solver.CountAreaOfDropletSurfaces,
		solver.CountAreaOfWaterAdjacentSurfaces,
	)
}
