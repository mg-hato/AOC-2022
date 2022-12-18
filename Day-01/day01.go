package main

import (
	"aoc/argshandle"
)

type CaloryList struct {
	list [][]int
}

func main() {
	argshandle.AoC2022DefaultProgram(
		ReadCaloryList,
		func(cl *CaloryList) int { return GetTotalCaloriesSumOfTopN(cl, 1) },
		func(cl *CaloryList) int { return GetTotalCaloriesSumOfTopN(cl, 3) },
	)
}
