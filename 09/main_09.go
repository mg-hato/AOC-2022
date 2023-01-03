package main

import (
	"aoc/argshandle"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(NewMotionSeriesReader),
		SimulateRopeWithTailCount(1), SimulateRopeWithTailCount(9),
	)
}
