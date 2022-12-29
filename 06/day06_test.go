package main

import (
	"aoc/testers"
	"testing"
)

func TestDay06_Reader(t *testing.T) {
	tester := testers.DefaultReaderTesterForComparableTypes(ReadDatastream, "ReadDatastream")

	tester.RunBothGoodAndErrorInputTests(t)

	tester.AddGoodInputTests(
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	)

	tester.AddErrorInputTests(
		"There is a white-space in the data stream",
		"There is an upper-case character in the data stream (concretely, \"X\")",
	)
}

func TestDay06_Solver(t *testing.T) {
	tester := testers.DefaultSolverTesterForComparableTypeR(
		FindFirstSequenceOfDifferent(4),
		FindFirstSequenceOfDifferent(14),
		"FindPacketStartMarker",
		"FindMessageStartMarker",
	)

	tester.AddTest("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19)
	tester.AddTest("bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23)
	tester.AddTest("nppdvjthqldpwncqszvftbrmjlhg", 6, 23)
	tester.AddTest("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29)
	tester.AddTest("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26)

	tester.RunBothSolversTests(t)
}
