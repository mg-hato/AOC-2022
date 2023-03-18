package main

import (
	"aoc/functional"
	"aoc/testers"
	"testing"
)

// Quick way to make a round structure
func r(abc ABC, xyz XYZ) Round {
	return Round{abc, xyz}
}

// Quick way to make an encrypted strategy guide
func esg(rounds ...Round) EncryptedStrategyGuide {
	return EncryptedStrategyGuide{rounds: rounds}
}

func TestDay02_Reader(t *testing.T) {
	tester := testers.DefaultReaderTester(ReadStrategy, "ReadEncryptedStrategyGuide")
	tester.ProvideEqualityFunctionForTypeT(areEqual)
	tester.AddGoodInputTests(
		esg(r(A, Y), r(B, X), r(C, Z)),
		esg(r(C, X)),
		esg(r(A, Z), r(A, X), r(B, X), r(C, Y), r(C, X), r(C, X)),
	)

	tester.AddErrorInputTests(
		"Line 1 has the ABC and XYZ switched up",
		"Line 3 is empty (not allowed)",
	)

	tester.RunBothGoodAndErrorInputTests(t)
}

func areEqual(lhs, rhs EncryptedStrategyGuide) bool {
	return functional.ArrayEqual(lhs.rounds, rhs.rounds)
}

func TestDay02_Solver(t *testing.T) {

	// A reminder of scoring: Shapes[R -> 1, P -> 2, S -> 3] & Outcome[W -> 6, Draw -> 3, L -> 0]

	tester := testers.DefaultSolverTesterForComparableTypeR(
		func(esg EncryptedStrategyGuide) (int, error) { return CalculateScore(esg, DirectlyAsShape) },
		func(esg EncryptedStrategyGuide) (int, error) { return CalculateScore(esg, AsDesiredOutcome) },
		"CalculateScore_XYZ_IsAShape",
		"CalculateScore_XYZ_IsDesiredOutcome",
	)

	// Shape values: Rock, Paper, Scissors => 1, 2, 3 (in that order)
	// Outcome values: Lose, Draw, Win => 0, 3, 6 (in that order)

	// Part 1: X, Y, Z => Rock, Paper, Scissors (in that order)
	// Part 2: X, Y, Z => Lose, Draw, Win (in that order)

	tester.AddTest(esg(r(A, X)), 4, 3) // RR => 1+3; Lose => RS => 3+0
	tester.AddTest(esg(r(A, Z)), 3, 8) // RS => 3+0; Win => RP => 2+6
	tester.AddTest(esg(r(C, X)), 7, 2) // SR => 1+6; Lose => SP => 2+0
	tester.AddTest(esg(r(B, Y)), 5, 5) // PP => 2+3; Draw => PP => 2+3
	tester.AddTest(esg(r(C, Z)), 6, 7) // SS => 3+3; Win => SR => 1+6
	tester.AddTest(esg(r(B, Z)), 9, 9) // PS => 3+6; Win => PS => 3+6

	tester.RunBothSolversTests(t)
}
