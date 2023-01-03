package main

import (
	. "aoc/functional"
	"aoc/reading"
	"aoc/testers"
	"testing"
)

func TestReader(t *testing.T) {
	tester := testers.DefaultReaderTester(
		reading.ReadWith(NewMotionSeriesReader),
		"ReadMotionSeries",
	)

	tester.ProvideEqualityFunctionForTypeT(ArrayEqual[Motion])

	tester.AddGoodInputTests(
		[]Motion{
			motion("L", 2),
			motion("R", 2),
			motion("R", 8),
			motion("U", 10),
			motion("D", 16),
		},
		[]Motion{
			motion("D", 1001),
			motion("R", 50),
			motion("U", 49),
			motion("L", 1),
		},
	)

	tester.AddErrorInputTests(
		"Line 3 has \"U68\" (they need to be separated by white-space)",
		"Line 2 has only \"R\" (needs a number after it)",
	)

	tester.RunBothGoodAndErrorInputTests(t)
}
