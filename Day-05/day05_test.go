package main

import (
	. "aoc/functional"
	"aoc/testers"
	"testing"
)

// quick-stacks creator
func s(stacks ...string) []string {
	return stacks
}

// quick-move creator
func m(qty, from, to int) Move {
	return Move{qty: qty, source: from, destination: to}
}

// quick-moves creator
func ms(moves ...Move) []Move {
	return moves
}

func TestDay05_Reader(t *testing.T) {
	tester := testers.DefaultReaderTester(ReadRearrangementPlan, "ReadRearrangementPlan")
	tester.ProvideEqualityFunctionForTypeT(equalsForRPlan)

	tester.AddGoodInputTests(
		RearrangementPlan{s("HX", "BJHLSF", "RMDHJTQ", "SGRHZ"), ms(m(1, 3, 1), m(1, 4, 2), m(2, 3, 1))},
		RearrangementPlan{s("A", "B", "C"), ms(m(1, 1, 2), m(1, 3, 1), m(1, 2, 3))},
	)

	tester.AddErrorInputTests(
		"Stack #3 has a floating box",
		"Stack #4 has a box \"3\" (only upper-case box names allowed)",
		"Stack #2 has a box \"j\" (only upper-case box names allowed)",
		"Stack ids should be (in order): 1,2,3,4",
		"Stacks are not aligned properly", // #5
		"Move instruction #2 moves from non-existing stack-id: stack-id 7",
		"Move instruction #3 moves to non-existing stack-id: stack-id 0",
		"Move instruction #1 typo: \"moev\" instead of \"move\"",
	)

	tester.RunBothGoodAndErrorInputTests(t)
}

func equalsForRPlan(lhs, rhs RearrangementPlan) bool {
	return ArrayEqual(lhs.stacks, rhs.stacks) && ArrayEqual(lhs.moves, rhs.moves)
}

func TestDay05_Solver(t *testing.T) {
	tester := testers.DefaultSolverTesterForComparableTypeR(
		func(plan RearrangementPlan) string { return FollowPlan(plan, CrateMover9000{}) },
		func(plan RearrangementPlan) string { return FollowPlan(plan, CrateMover9001{}) },
		"FollowPlanWithCrateMover9000", "FollowPlanWithCrateMover9001",
	)

	// Example input given in the problem statement
	tester.AddTest(RearrangementPlan{
		s("ZN", "MCD", "P"),
		ms(
			m(1, 2, 1),
			m(3, 1, 3),
			m(2, 2, 1),
			m(1, 1, 2),
		)}, "CMZ", "MCD")

	// Custom input #1
	tester.AddTest(RearrangementPlan{
		s("A", "B", "C"),
		ms(
			m(1, 1, 2),
			m(1, 3, 1),
			m(1, 2, 3),
		)}, "CBA", "CBA")

	// Custom input #2
	tester.AddTest(RearrangementPlan{
		s("ABC", "D", "E"),
		ms(
			// ABC D E     CM9000    | CM9001
			m(3, 1, 2), // . DCBA E  | . DABC E
			m(1, 3, 1), // E DCBA .  | E DABC .
			m(2, 2, 3), // E DC AB   | E DA BC
			m(1, 2, 1), // EC D AB   | EA D BC
		)}, "CDB", "ADC")
	tester.RunBothSolversTests(t)
}
