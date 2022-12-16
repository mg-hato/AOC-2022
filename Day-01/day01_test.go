package main

import (
	"fmt"
	"testing"
)

func TestDay01_TestReader(t *testing.T) {
	expected := map[int]CaloryList{
		1: {list: [][]int{{10, 20, 30}, {40, 50, 60}, {70}}},
		2: {list: [][]int{{12, 12}, {12, 12, 12}}},
		3: {list: [][]int{{10}}},
		4: {list: [][]int{}},
	}
	for k, exp := range expected {
		cl, _ := ReadCaloryList(fmt.Sprintf("./test/input.%d", k))
		if !cl.match(exp) {
			t.Errorf("Input #%d mismatch.\n\tRead: %v\n\tExpected: %v\n", k, *cl, exp)
		}
	}

}

func (lhs CaloryList) match(rhs CaloryList) bool {
	if len(lhs.list) != len(rhs.list) {
		return false
	}

	for i, list := range lhs.list {
		if len(list) != len(rhs.list[i]) {
			return false
		}

		for j, v := range list {
			if v != rhs.list[i][j] {
				return false
			}
		}
	}
	return true
}

func TestDay01_Solver(t *testing.T) {
	inputs := []CaloryList{
		{list: [][]int{{100}, {40, 50, 60}, {70}}},
		{list: [][]int{{100}, {40, 50, 60}, {70}}},
		{list: [][]int{}},
		{list: [][]int{{10}}},
	}

	n_values := []int{
		1,
		2,
		6,
		3,
	}

	expected := []int{
		150,
		250,
		0,
		10,
	}

	for i, input := range inputs {
		if expected[i] != GetTotalCaloriesSumOfTopN(&input, n_values[i]) {
			t.Errorf(
				"Solver test #%d mismatch. Expected: %d for n = %d and input %v",
				i, expected[i], n_values[i], input,
			)
		}
	}
}
