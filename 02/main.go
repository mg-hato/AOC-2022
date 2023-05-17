package main

import (
	"aoc/argshandle"
	"aoc/day02/reader"
	s "aoc/day02/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.StategyGuideReader),
		s.CalculateScore(s.ShapeBasedRoundInterpreter()),
		s.CalculateScore(s.OutcomeBasedRoundInterpreter()),
	)
}
