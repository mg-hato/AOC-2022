package main

import (
	"aoc/testers"
	"testing"
)

func TestSolver(t *testing.T) {
	tester := testers.DefaultSolverTesterForComparableTypeR(
		AnalyseForestWith(NewVisibilityAnalyser),
		AnalyseForestWith(NewScenicScoreAnalyser),
		"CountVisibleTrees",
		"GetBestScenicScore",
	)

	// Example given in the problem statement
	tester.AddTest([][]byte{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}, 21, 8)

	// Custom input #1
	// All trees are on the edge => all are visible
	tester.AddTest([][]byte{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
	}, 10, 0)

	// Custom input #2
	// Only tree of height 5 is not visible
	// Tree of height 8 is visible from down
	// All others are on edge
	tester.AddTest([][]byte{
		{9, 9, 9},
		{9, 5, 9},
		{9, 8, 9},
		{9, 7, 9},
	}, 11, 2)

	tester.RunBothSolversTests(t)
}
