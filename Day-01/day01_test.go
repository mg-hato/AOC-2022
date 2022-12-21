package main

import (
	"aoc/functional"
	"fmt"
	"testing"
)

func TestDay01_TestReader(t *testing.T) {
	expected := map[int]List{
		1: {calories: [][]int{{10, 20, 30}, {40, 50, 60}, {70}}},
		2: {calories: [][]int{{12, 12}, {12, 12, 12}}},
		3: {calories: [][]int{{10}}},
		4: {calories: [][]int{}},
	}
	for k, exp := range expected {
		filename := fmt.Sprintf("./test/input.%d", k)
		if readList, _ := ReadList(filename); !readList.equals(exp) {
			t.Errorf("Test #%d failed: ReadList(\"%s\")", k, filename)
			t.Errorf("\tReturned: %v", readList)
			t.Errorf("\tExpected: %v", exp)
		}
	}

}

func (lhs List) match(rhs List) bool {
	if len(lhs.calories) != len(rhs.calories) {
		return false
	}

	for i, list := range lhs.calories {
		if len(list) != len(rhs.calories[i]) {
			return false
		}

		for j, v := range list {
			if v != rhs.calories[i][j] {
				return false
			}
		}
	}
	return true
}

func TestDay01_Solver(t *testing.T) {
	tests := []struct {
		input       List
		n, expected int
	}{
		{
			input: List{[][]int{}},
			n:     10, expected: 0,
		},
		{
			input: List{[][]int{{100}, {40, 50, 60}, {70}}},
			n:     1, expected: 150,
		},
		{
			input: List{[][]int{{100}, {40, 50, 60}, {70}}},
			n:     2, expected: 250,
		},
		{
			input: List{[][]int{{100}, {40, 50, 60}, {70}}},
			n:     3, expected: 320,
		},
		{
			// sums => 1, 9, 5, 6, 7, 17, 10
			// top 4 sums => 17, 10, 9, 7
			input: List{[][]int{{1}, {2, 3, 4}, {5}, {6}, {7}, {8, 9}, {10}}},
			n:     4, expected: 43,
		},
	}

	for tn, test := range tests {
		if result := GetTotalCaloriesSumOfTopN(&test.input, test.n); result != test.expected {
			t.Errorf("Test #%d failed: GetTotalCaloriesSumOfTOpN(%v, %d)", tn+1, test.input, test.n)
			t.Errorf("\tReturned: %d", result)
			t.Errorf("\tExpected: %d", test.expected)

		}
	}
}

// Tests whether to List structs are equal
func (lhs List) equals(rhs List) bool {
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
