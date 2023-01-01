package main

import (
	"aoc/testers"
	"testing"

	. "aoc/functional"
)

func TestVisibilityAnalyser(t *testing.T) {
	tester := testers.DefaultTester(
		func(heights []byte) []bool {
			trees := Map(func(p Pair[int, byte]) Tree { return Tree{p.First, p.Second} }, Enumerate(heights))
			analyser := VisibilityAnalyster{visibility: map[int]bool{}}
			analyser.AnalyseForestRow(trees)

			return Map(func(t Tree) bool { return analyser.visibility[t.id] }, trees)
		}, "GetVisibilities",
	)

	tester.ProvideEqualityFunction(ArrayEqual[bool])

	tester.AddTest([]byte{1, 2, 3, 4, 5}, []bool{true, true, true, true, true})
	tester.AddTest([]byte{5, 4, 1, 5}, []bool{true, false, false, false})
	tester.AddTest([]byte{3, 2, 3, 5}, []bool{true, false, false, true})

	tester.RunTests(t)
}

func TestScenicScoreAnalyser(t *testing.T) {
	tester := testers.DefaultTester(
		func(heights []byte) []int {
			trees := Map(func(p Pair[int, byte]) Tree { return Tree{p.First, p.Second} }, Enumerate(heights))
			analyser := ScenicScoreAnalyser{scenic_scores: map[int]int{}}
			analyser.AnalyseForestRow(trees)

			return Map(func(t Tree) int { return analyser.scenic_scores[t.id] }, trees)
		}, "GetScenicScores",
	)

	tester.ProvideEqualityFunction(ArrayEqual[int])

	tester.AddTest([]byte{1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4})
	tester.AddTest([]byte{5, 4, 1, 5}, []int{0, 1, 1, 3})
	tester.AddTest([]byte{3, 2, 3, 5}, []int{0, 1, 2, 3})

	tester.RunTests(t)
}
