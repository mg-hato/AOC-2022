package main

import (
	"aoc/argshandle"
	"aoc/d17/reader"
	"aoc/d17/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.JetPatternReader),
		solver.GetHeightAfterNRocks(2022),
		solver.GetHeightAfterNRocks(1_000_000_000_000),
	)
}

// 1 000 000 000 000
