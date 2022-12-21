package functional

import (
	"fmt"
	"strings"
	"testing"
)

func TestForEach(t *testing.T) {

	tests := []struct {
		input                          []int
		expectedSum, expectedCallCount int
	}{
		{input: []int{}, expectedSum: 0, expectedCallCount: 0},
		{input: []int{10, 20}, expectedSum: 30, expectedCallCount: 2},
		{input: []int{100, 70, 1000}, expectedSum: 1200, expectedCallCount: 5},
	}

	var sum, callCount int = 0, 0
	f := func(x int) {
		callCount++
		sum += x
	}

	for test_number, test := range tests {
		ForEach(f, test.input)
		if sum != test.expectedSum || callCount != test.expectedCallCount {
			msg := strings.Join([]string{
				fmt.Sprintf("Test %d failed on ForEach(f, %v). Call count and/or total sum do not match!\n", test_number, test.input),
				fmt.Sprintf("\tActual call count: %d\n", callCount),
				fmt.Sprintf("\tExpected call count: %d\n", test.expectedCallCount),
				fmt.Sprintf("\tActual total sum: %d\n", sum),
				fmt.Sprintf("\tExpected total sum: %d\n", test.expectedSum),
			}, "")
			t.Error(msg)
		}
	}
}
