package main

import (
	"aoc/argshandle"
	"aoc/day08/reader"
	s "aoc/day08/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.ForestReader),
		s.AnalyseForest(s.VisibilityTreeLineAnalyser),
		s.AnalyseForest(s.ScenicScoreTreeLineAnalyser),
	)
}
