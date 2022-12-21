package functional

import (
	"fmt"
	"testing"
)

func TestMap_Double(t *testing.T) {
	tests := []struct {
		input, expectedOutput []int
	}{
		{input: []int{1, 2, 3}, expectedOutput: []int{2, 4, 6}},
		{input: []int{}, expectedOutput: []int{}},
		{input: []int{1000, -5}, expectedOutput: []int{2000, -10}},
		{input: []int{55}, expectedOutput: []int{110}},
	}

	doublerFunc := func(i int) int { return i * 2 }

	for test_number, test := range tests {
		if result := Map(doublerFunc, test.input); !ArrayEqual(result, test.expectedOutput) {
			t.Errorf("Test #%d failed. Map(doublerFunc, %v) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}

func TestMap_AddNumber(t *testing.T) {

	tests := []struct {
		input, expectedOutput []int
	}{
		{input: []int{1, 2, 3}, expectedOutput: []int{8, 9, 10}},
		{input: []int{}, expectedOutput: []int{}},
		{input: []int{-3, 5, 100, 33}, expectedOutput: []int{4, 12, 107, 40}},
	}

	add7 := func(i int) int { return i + 7 }

	for test_number, test := range tests {
		if result := Map(add7, test.input); !ArrayEqual(result, test.expectedOutput) {
			t.Errorf("Test #%d failed. Map(addNumber, %v) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}

func TestMap_ConcatNumbers(t *testing.T) {

	tests := []struct {
		input          []int
		expectedOutput []string
	}{
		{input: []int{0, 100, 55, 7}, expectedOutput: []string{"00", "100100", "5555", "77"}},
		{input: []int{}, expectedOutput: []string{}},
		{input: []int{9}, expectedOutput: []string{"99"}},
	}

	printTwice := func(i int) string { return fmt.Sprintf("%d%d", i, i) }

	for test_number, test := range tests {
		if result := Map(printTwice, test.input); !ArrayEqual(result, test.expectedOutput) {
			t.Errorf("Test #%d failed. Map(printTwice, %v) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}

}
