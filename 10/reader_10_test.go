package main

import (
	. "aoc/functional"
	"aoc/reading"
	"aoc/testers"
	"testing"
)

func TestReader10(t *testing.T) {
	tester := testers.DefaultReaderTester(
		reading.ReadWith(NewInstructionReader),
		"ReadInstructions",
	)

	tester.ProvideEqualityFunctionForTypeT(ArrayEqualWith(func(lhs, rhs Instruction) bool { return lhs.String() == rhs.String() }))

	tester.AddGoodInputTests(
		[]Instruction{NewNoop(), NewAddx(3), NewAddx(-5)},
		[]Instruction{NewAddx(12), NewAddx(20), NewNoop(), NewNoop(), NewAddx(0), NewNoop(), NewAddx(-101)},
	)

	tester.AddErrorInputTests(
		"Line 3: noop instruction takes no arguments",
		"Line 4: addx instruction takes exactly one argument",
		"Line 2: addx instruction takes exactly one argument",
		"Line 4: typo \"add\"",
	)

	tester.RunBothGoodAndErrorInputTests(t)
}
