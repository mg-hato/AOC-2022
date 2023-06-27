package main

import (
	"aoc/argshandle"
	"aoc/d19/reader"
	"aoc/d19/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.BlueprintsReader),
		solver.AnalyseBlueprintQualities(24),
		solver.AnalyseBlueprintQualityOfFirstN(3, 32),
	)
}
