package main

import (
	"aoc/argshandle"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		ReadDatastream,
		FindFirstSequenceOfDifferent(4),
		FindFirstSequenceOfDifferent(14),
	)
}
