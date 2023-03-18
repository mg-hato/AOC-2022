package main

import (
	"aoc/reading"
	"aoc/testers"
	"strings"
	"testing"
)

func TestSolver10(t *testing.T) {

	// Because example input is quite long,
	// in this test we read it from a file & call the solver

	withSolver := func(solver func([]Instruction) (Result, error)) func(string) (Result, error) {
		return func(input_file string) (Result, error) {
			instructions, err := reading.ReadWith(NewInstructionReader)(input_file)
			if err != nil {
				t.Errorf("Input could not be read from file \"%s\"", input_file)
				return nil, nil
			} else {
				return solver(instructions)
			}
		}
	}

	tester := testers.DefaultSolverTester(
		withSolver(SignalStrengths),
		withSolver(DrawCRT),
		"GetSignalStrengths",
		"DrawCRT",
	)

	tester.ProvideEqualityFunctionForTypeR(func(lhs, rhs Result) bool { return lhs == rhs })

	var result_crt string = strings.Join(
		[]string{
			"##..##..##..##..##..##..##..##..##..##..",
			"###...###...###...###...###...###...###.",
			"####....####....####....####....####....",
			"#####.....#####.....#####.....#####.....",
			"######......######......######......####",
			"#######.......#######.......#######.....",
		}, "")

	tester.AddTest("./test/example.txt", ResultInt{13140}, ResultCRT{result_crt})
	tester.RunBothSolversTests(t)
}
