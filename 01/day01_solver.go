package main

import (
	"sort"

	. "aoc/functional"
)

func GetTotalCaloriesSumOfTopN(list List, n int) (int, error) {
	var sums []int = Map(Sum[int], list.calories)
	sort.Slice(sums, func(i, j int) bool { return sums[i] > sums[j] })
	return Sum(Take(n, sums)), nil
}
