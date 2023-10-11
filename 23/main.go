package main

import (
	"aoc/argshandle"
	"aoc/d23/reader"
	"aoc/d23/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.GroveMapReader),
		solver.CountFreeSpacesInEncapsulatingRegion(10),
		solver.FirstRoundWhenNoElfMoves(10_000),
	)
}
