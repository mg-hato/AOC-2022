package main

import (
	"aoc/argshandle"
	"aoc/d24/reader"
	"aoc/d24/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.ValleyMapReader),
		solver.CalculateSingleTripBestTime,
		solver.CalculateTripleTripBestTime,
	)
}
