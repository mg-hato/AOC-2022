package functional

import (
	"fmt"
	"testing"
)

func TestFoldl_ConcatNumbersInReverse(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput string
	}{
		{input: []int{1, 2, 3}, expectedOutput: "321"},
		{input: []int{}, expectedOutput: ""},
		{input: []int{1000, 5}, expectedOutput: "51000"},
		{input: []int{55}, expectedOutput: "55"},
	}

	concatNumberToString := func(str string, i int) string { return fmt.Sprintf("%d%s", i, str) }

	for test_number, test := range tests {
		if result := Foldl(concatNumberToString, test.input, ""); result != test.expectedOutput {
			t.Errorf("Test #%d failed. Foldl(concatNumberToString, %v, \"\") returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}

func TestFoldl_Product(t *testing.T) {
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
		if result := Foldl(multiply, test.input, 1); result != test.expectedOutput {
			t.Errorf("Test #%d failed. Foldl(multiply, %v, 1) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}

func TestFoldl_SumPrint(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput string
	}{
		{input: []int{1, 2, 3}, expectedOutput: "(((X + 1) + 2) + 3)"},
		{input: []int{}, expectedOutput: "X"},
		{input: []int{1000, 5}, expectedOutput: "((X + 1000) + 5)"},
		{input: []int{55}, expectedOutput: "(X + 55)"},
	}

	sumPrint := func(str string, i int) string { return fmt.Sprintf("(%s + %d)", str, i) }

	for test_number, test := range tests {
		if result := Foldl(sumPrint, test.input, "X"); result != test.expectedOutput {
			t.Errorf("Test #%d failed. Foldl(sumPrint, %v, \"\") returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}
