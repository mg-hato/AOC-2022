package main

import (
	"aoc/argshandle"
	"aoc/d21/reader"
	"aoc/d21/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.MonkeyJobsReader),
		solver.SolveRoot,
		solver.SolveHumn,
	)
}
