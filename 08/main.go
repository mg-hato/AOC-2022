package main

import (
	"aoc/argshandle"
	"aoc/d08/reader"
	s "aoc/d08/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.ForestReader),
		s.AnalyseForest(s.VisibilityTreeLineAnalyser),
		s.AnalyseForest(s.ScenicScoreTreeLineAnalyser),
	)
}
