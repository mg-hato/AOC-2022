package main

import (
	"aoc/argshandle"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		ReadList,
		func(cl List) int { return GetTotalCaloriesSumOfTopN(cl, 1) },
		func(cl List) int { return GetTotalCaloriesSumOfTopN(cl, 3) },
	)
}
