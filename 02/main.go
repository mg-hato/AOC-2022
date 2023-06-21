package main

import (
	"aoc/argshandle"
	"aoc/d02/reader"
	s "aoc/d02/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.StategyGuideReader),
		s.CalculateScore(s.ShapeBasedRoundInterpreter()),
		s.CalculateScore(s.OutcomeBasedRoundInterpreter()),
	)
}
