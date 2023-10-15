package main

import (
	"aoc/argshandle"
	"aoc/d25/reader"
	"aoc/d25/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.SnafuNumbersReader),
		solver.SolveSnafu,
		solver.SolveSnafu,
	)
}
