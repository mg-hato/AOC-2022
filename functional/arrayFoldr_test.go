package functional

import (
	"fmt"
	"testing"
)

func TestFoldr_ConcatNumbers(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput string
	}{
		{input: []int{1, 2, 3}, expectedOutput: "123"},
		{input: []int{}, expectedOutput: ""},
		{input: []int{1000, 5}, expectedOutput: "10005"},
		{input: []int{55}, expectedOutput: "55"},
	}

	concatNumberToString := func(i int, str string) string { return fmt.Sprintf("%d%s", i, str) }

	for test_number, test := range tests {
		if result := Foldr(concatNumberToString, test.input, ""); result != test.expectedOutput {
			t.Errorf("Test #%d failed. Foldr(concatNumberToString, %v, \"\") returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}

func TestFoldr_Product(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput int
	}{
		{input: []int{1, 2, 3}, expectedOutput: 6},
		{input: []int{}, expectedOutput: 1},
		{input: []int{1000, 5}, expectedOutput: 5000},
		{input: []int{55}, expectedOutput: 55},
		{input: []int{10, 20, 0, 30}, expectedOutput: 0},
	}

	multiply := func(x int, y int) int { return x * y }

	for test_number, test := range tests {
		if result := Foldr(multiply, test.input, 1); result != test.expectedOutput {
			t.Errorf("Test #%d failed. Foldr(multiply, %v, 1) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}

func TestFoldr_SumPrint(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput string
	}{
		{input: []int{1, 2, 3}, expectedOutput: "(1 + (2 + (3 + X)))"},
		{input: []int{}, expectedOutput: "X"},
		{input: []int{1000, 5}, expectedOutput: "(1000 + (5 + X))"},
		{input: []int{55}, expectedOutput: "(55 + X)"},
	}

	sumPrint := func(i int, str string) string { return fmt.Sprintf("(%d + %s)", i, str) }

	for test_number, test := range tests {
		if result := Foldr(sumPrint, test.input, "X"); result != test.expectedOutput {
			t.Errorf("Test #%d failed. Foldr(sumPrint, %v, \"\") returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}
