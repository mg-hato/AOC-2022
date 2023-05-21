package main

import (
	"aoc/argshandle"
	"aoc/day06/reader"
	"aoc/day06/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.DatastreamBufferReader),
		solver.FindPositionOfTheFirstMarker(4),
		solver.FindPositionOfTheFirstMarker(14),
	)
}
