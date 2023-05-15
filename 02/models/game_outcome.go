package models

// Outcome of one rock-paper-scissors game
type GameOutcome int

const (
	Lose GameOutcome = iota
	Draw
	Win
)
