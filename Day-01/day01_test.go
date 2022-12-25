package main

import (
	"aoc/functional"
	"aoc/testers"
	"testing"
)

func TestDay01_Reader(t *testing.T) {
	tester := testers.DefaultReaderTester(
		ReadList,
		"ReadListOfCalories",
	)

	tester.ProvideEqualityFunctionForTypeT(areEqual)

	tester.AddGoodInputTests(
		List{[][]int{{10, 20, 30}, {40, 50, 60}, {70}}},
		List{[][]int{{12, 12}, {12, 12, 12}}},
		List{[][]int{{10}}},
		List{[][]int{}},
	)

	tester.AddErrorInputTests("Line 6 has number with commas (not accepted)")

	tester.RunBothGoodAndErrorInputTests(t)
}

func TestDay01_Solver(t *testing.T) {
	tester := testers.DefaultSolverTesterForComparableTypeR(
		func(list List) int { return GetTotalCaloriesSumOfTopN(list, 1) },
		func(list List) int { return GetTotalCaloriesSumOfTopN(list, 3) },
		"GetSumOfTop1",
		"GetSumOfTop3",
	)

	tester.AddTest(List{[][]int{{1}, {2}, {3}}}, 3, 6)
	tester.AddTest(List{[][]int{{1, 2, 3, 4}, {81}, {5, 5, 5}, {71}}}, 81, 167)
	tester.AddTest(List{[][]int{{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}}}, 1, 3)
	tester.AddTest(List{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}}, 3, 9)

	tester.RunBothSolversTests(t)
}

func areEqual(lhs, rhs List) bool {
	size := len(lhs.calories)
	if size != len(rhs.calories) {
		return false
	}

	for i := 0; i < size; i++ {
		if !functional.ArrayEqual(lhs.calories[i], rhs.calories[i]) {
			return false
		}
	}
	return true
}
