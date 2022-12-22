package main

import "aoc/argshandle"

func main() {
	argshandle.AoC2022DefaultProgram(
		ReadStrategy,
		func(s *EncryptedStrategyGuide) int { return CalculateScore(s, DirectlyAsShape) },
		func(s *EncryptedStrategyGuide) int { return CalculateScore(s, AsDesiredOutcome) },
	)

}
