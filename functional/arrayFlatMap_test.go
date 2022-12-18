package functional

import (
	"fmt"
	"testing"
)

func TestFlatMap_Digits(t *testing.T) {
	tests := []struct {
		input, expectedOutput []int
	}{
		{input: []int{}, expectedOutput: []int{}},
		{input: []int{10, 123}, expectedOutput: []int{1, 0, 1, 2, 3}},
		{input: []int{14560}, expectedOutput: []int{1, 4, 5, 6, 0}},
	}

	getDigits := func(x int) []int {
		s := fmt.Sprint(x)
		digits := make([]int, len(s))
		for i, c := range fmt.Sprint(x) {
			digits[i] = int(c - '0')
		}
		return digits
	}

	for test_number, test := range tests {
		if result := FlatMap(getDigits, test.input); !ArrayEqual(result, test.expectedOutput) {
			t.Errorf("Test #%d failed. FlatMap(getDigits, %v) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}

func TestFlatMap_factorisation(t *testing.T) {
	tests := []struct {
		input, expectedOutput []int
	}{
		{input: []int{}, expectedOutput: []int{}},
		{input: []int{10, 20}, expectedOutput: []int{2, 5, 2, 2, 5}},
		{input: []int{100, 77, 67}, expectedOutput: []int{2, 2, 5, 5, 7, 11, 67}},
	}

	factorise := func(x int) []int {
		var p int = 2
		var factors []int = make([]int, 0)
		for x != 1 {
			if x%p == 0 {
				factors = append(factors, p)
				x /= p
			} else {
				p++
			}
		}
		return factors
	}

	for test_number, test := range tests {
		if result := FlatMap(factorise, test.input); !ArrayEqual(result, test.expectedOutput) {
			t.Errorf("Test #%d failed. FlatMap(factorise, %v) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}
