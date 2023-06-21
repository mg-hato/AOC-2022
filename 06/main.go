package main

import (
	"aoc/argshandle"
	"aoc/d06/reader"
	"aoc/d06/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.DatastreamBufferReader),
		solver.FindPositionOfTheFirstMarker(4),
		solver.FindPositionOfTheFirstMarker(14),
	)
}
