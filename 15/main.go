package main

import (
	"aoc/argshandle"
	"aoc/d15/reader"
	"aoc/d15/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.SensorReportsReader),
		solver.BeaconExclusionCount(2_000_000),
		solver.DistressBeaconTuningFrequencyFinder(4_000_000),
	)
}
