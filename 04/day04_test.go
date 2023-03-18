package main

import (
	. "aoc/functional"
	"aoc/testers"
	"testing"
)

// quick-way to initialise assignment pair
func ap(a, b, c, d int) AssignmentPair {
	return AssignmentPair{
		first:  SectionRange{left: a, right: b},
		second: SectionRange{left: c, right: d},
	}
}

func TestDay04_Reader(t *testing.T) {
	tester := testers.DefaultReaderTester(ReadListOfAssignmentPairs, "ReadAssignmentPairs")
	tester.ProvideEqualityFunctionForTypeT(ArrayEqual[AssignmentPair])

	tester.AddGoodInputTests(
		[]AssignmentPair{ap(1, 5, 10, 10), ap(711, 1000, 1, 1), ap(3, 4, 5, 5)},
		[]AssignmentPair{ap(399, 401, 400, 400)},
		[]AssignmentPair{},
	)

	tester.AddErrorInputTests(
		"Line 3 defines too many ranges",
		"Line 4 defines invalid range (left bigger than right)",
		"Line 4 range defined poorly with not a number",
	)

	tester.RunBothGoodAndErrorInputTests(t)
}

func TestDay04_Solver(t *testing.T) {
	tester := testers.DefaultSolverTesterForComparableTypeR(
		func(pairs []AssignmentPair) (int, error) {
			return CountAssignmentPairsThatSatisfy(OneFullyContainsTheOther, pairs)
		},
		func(pairs []AssignmentPair) (int, error) {
			return CountAssignmentPairsThatSatisfy(SectionRangesOverlap, pairs)
		},
		"CountAssignmentPairsWhereOneFullyContainsTheOther",
		"CountAssignmentPairsThatOverlap",
	)

	// Example from the problem statement
	tester.AddTest(
		[]AssignmentPair{
			ap(2, 4, 6, 8),
			ap(2, 3, 4, 5),
			ap(5, 7, 7, 9),
			ap(2, 8, 3, 7),
			ap(6, 6, 4, 6),
			ap(2, 6, 4, 8),
		}, 2, 4,
	)

	// Custom test #1
	tester.AddTest(
		[]AssignmentPair{
			ap(1, 1, 1, 2),         // F O
			ap(1, 1, 1, 2),         // F O
			ap(1, 5, 7, 9),         //
			ap(100, 200, 150, 250), // O
			ap(100, 200, 130, 201), // O
			ap(100, 200, 201, 201), //
			ap(51, 60, 55, 60),     // F O
			ap(2, 8, 3, 7),         // F O
			ap(6, 6, 4, 6),         // F O
			ap(2, 6, 4, 8),         // O
			ap(15, 167, 10, 200),   // F O
			ap(15, 167, 168, 200),  //
		}, 6, 9,
	)

	tester.RunBothSolversTests(t)
}
