package solver

import m "aoc/d02/models"

// An interpreter for a single round / game of rock-paper-scissors from the strategy guide
type RoundInterpreter interface {
	String() string
	GetScore(m.Round) int
}
