package main

import (
	"fmt"
)

// One of the shapes: Rock, Paper or Scissors
type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

// Outcome of the game: Win, Lose or Draw
type Outcome int

const (
	Draw Outcome = iota
	Win
	Lose
)

// Enumeration used by opponent player: A, B or C
type ABC int

const (
	A ABC = iota
	B
	C
)

// Print ABC enumeration as expected
func (abc ABC) String() string {
	switch abc {
	case A:
		return "A"
	case B:
		return "B"
	default: // must be C
		return "C"
	}
}

// Enumeration used for myself: X, Y or Z
type XYZ int

const (
	X XYZ = iota
	Y
	Z
)

// Print XYZ enumeration as expected
func (xyz XYZ) String() string {
	switch xyz {
	case X:
		return "X"
	case Y:
		return "Y"
	default: // must be Z
		return "Z"
	}
}

// Decoding strategy for "ABC" enumeration (Day 02: first & second task)
type XYZDecoderStrategy int

const (
	DirectlyAsShape XYZDecoderStrategy = iota
	AsDesiredOutcome
)

func (decStrat XYZDecoderStrategy) String() string {
	switch decStrat {
	case DirectlyAsShape:
		return "XYZ -> Shape"
	default: // Must be "AsDesiredOutcome"
		return "XYZ -> DesiredOutcome -> Shape"
	}
}

// Round represents one play of rock-paper-scissors game. Opponent's play is descibed with ABC enumeration,
// and my play is described with XYZ enumeration
type Round struct {
	opponent ABC
	me       XYZ
}

// An encrypted strategy guide is a list of "predicted" opponent plays (in ABC enumeration)
// and plays I should use to combate them (in "XYZ" enumeration)
type EncryptedStrategyGuide struct {
	rounds []Round
}

// Score obtained by choosing the provided shape (regardless of your oppoent's choice)
func (shape Shape) ShapeScore() int {
	switch shape {
	case Rock:
		return 1
	case Paper:
		return 2
	default: // Must be Scissors
		return 3
	}
}

// Score obtained based on the outcome of the game (Win, Lose or Draw)
func (outcome Outcome) OutcomeScore() int {
	switch outcome {
	case Lose:
		return 0
	case Draw:
		return 3
	default: // Must be Win
		return 6
	}
}

// Score obtained for the round (based on XYZ decoding strategy)
func (r Round) ScoreForRound(decodingStategyForABC XYZDecoderStrategy) int {

	// Get opponent's shape
	opponentShape := r.DecodeOpponent()

	// Get my shape (based on the decoding strategy)
	myShape := r.DecodeMe(decodingStategyForABC)

	shapeToOutcome := map[Shape]Outcome{
		WinAgainst(opponentShape):  Win,
		LoseAgainst(opponentShape): Lose,
		DrawAgainst(opponentShape): Draw,
	}

	return myShape.ShapeScore() + shapeToOutcome[myShape].OutcomeScore()
}

// Decode opponent's shape. Direct mapping: A -> Rock, B -> Paper, C -> Scissors
func (r Round) DecodeOpponent() Shape {
	switch r.opponent {
	case A:
		return Rock
	case B:
		return Paper
	default: // Must be C
		return Scissors
	}
}

// Decode my shape using the decoder strategy provided
func (r Round) DecodeMe(decodingStrategy XYZDecoderStrategy) Shape {
	switch decodingStrategy {
	case DirectlyAsShape: // Decode XYZ directly as one of the Rock, Paper or Scissors
		return r.DecodeXYZDirectly()
	default: // Decode XYZ as "desired outcome": Win, Lose or Draw
		return r.DecodeXYZBasedOnOutcome()
	}
}

// Deocde XYZ directly as one of the Rock, Paper or Scissors
func (r Round) DecodeXYZDirectly() Shape {
	switch r.me {
	case X:
		return Rock
	case Y:
		return Paper
	default: // Must be Z
		return Scissors
	}
}

// Decode my shape so that my shape yields the desired outcome based on XYZ enumeration
func (r Round) DecodeXYZBasedOnOutcome() Shape {
	switch r.me.DecodeAsOutcome() {
	case Win:
		return WinAgainst(r.DecodeOpponent())
	case Lose:
		return LoseAgainst(r.DecodeOpponent())
	default: // Must be "Draw"
		return DrawAgainst(r.DecodeOpponent())
	}
}

// Decode XYZ enumeration as desired outcome: X -> Lose, Y -> Draw, Z -> Win
func (xyz XYZ) DecodeAsOutcome() Outcome {
	switch xyz {
	case X:
		return Lose
	case Y:
		return Draw
	default: // Must be Z
		return Win
	}
}

// Returns a shape that wins against the provided shape
func WinAgainst(shape Shape) Shape {
	switch shape {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	default: // Must be Scissors, against them Rock wins
		return Rock
	}
}

// Returns a shape that loses against provided shape
func LoseAgainst(shape Shape) Shape {
	switch shape {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	default: // Must be Scissors, against them Paper loses
		return Paper
	}
}

// Returns a shape that will result in a draw against the provided shape
func DrawAgainst(shape Shape) Shape {
	return shape
}

func (r Round) String() string {
	return fmt.Sprintf("(%s:%s)", r.opponent, r.me)
}

func (s EncryptedStrategyGuide) String() string {
	return fmt.Sprintf("Strategy{ %s }", s.rounds)
}
