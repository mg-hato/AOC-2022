package main

import (
	"sort"
)

func GetTotalCaloriesSumOfTopN(caloryList *CaloryList, n int) int {
	var caloryCounts []int = caloryList.getCalorySumsPerElf()
	sort.Ints(caloryCounts)
	return sum(caloryCounts[max(len(caloryCounts)-n, 0):])
}

func (cl *CaloryList) getCalorySumsPerElf() []int {
	var sums []int = make([]int, len(cl.list))
	for i, calories := range cl.list {
		sums[i] = sum(calories)
	}
	return sums
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	} else {
		return rhs
	}
}
