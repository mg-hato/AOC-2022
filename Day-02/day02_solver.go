package main

import (
	. "aoc/functional"
)

// Solve Day 02 problem given the input and the XYZ decoding strategy (i.e. part 1 or part 2)
func CalculateScore(strategy *EncryptedStrategyGuide, xyzDecodeStrategy XYZDecoderStrategy) int {
	return Sum(Map(
		func(r Round) int { return r.ScoreForRound(xyzDecodeStrategy) },
		strategy.rounds,
	))
}
