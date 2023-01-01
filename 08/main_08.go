package main

import (
	"aoc/argshandle"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(NewForestReader),
		AnalyseForestWith(NewVisibilityAnalyser),
		AnalyseForestWith(NewScenicScoreAnalyser),
	)
}
