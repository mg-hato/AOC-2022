package main

import (
	"aoc/argshandle"
	"fmt"
	"os"
)

type CaloryList struct {
	list [][]int
}

func main() {

	ch := make(chan int, 1)

	argshandle.HandleArgumentsAndExecute(
		os.Args,
		ReadCaloryList,
		func(cl *CaloryList) int { return GetTotalCaloriesSumOfTopN(cl, 1) },
		func(cl *CaloryList) int { return GetTotalCaloriesSumOfTopN(cl, 3) },
		ch,
	)

	fmt.Println(<-ch)
}
