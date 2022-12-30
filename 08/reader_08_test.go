package main

import (
	. "aoc/functional"
	"aoc/reading"
	"aoc/testers"
	"testing"
)

func TestReader(t *testing.T) {
	tester := testers.DefaultReaderTester(
		reading.ReadWith(NewForestReader),
		"ReadForestGrid",
	)

	tester.ProvideEqualityFunctionForTypeT(ArrayEqualWith(ArrayEqual[byte]))

	tester.AddGoodInputTests(
		[][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]byte{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
	)

	tester.AddErrorInputTests(
		"Row 2 has four trees, but other rows have three",
		"Row 3 has non-digit character (character 'o')",
	)

	tester.RunBothGoodAndErrorInputTests(t)
}
