package functional

import (
	"fmt"
	"strings"
	"testing"
)

func TestTake_test(t *testing.T) {
	tests := []struct {
		n               int
		input, expected []int
	}{
		{n: 3, input: []int{1}, expected: []int{1}},
		{n: 20, input: []int{1}, expected: []int{1}},
		{n: 0, input: []int{1}, expected: []int{}},
		{n: 3, input: []int{10, 15, 20, 100, 1010, 7050}, expected: []int{10, 15, 20}},
		{n: 1, input: []int{10, 15, 20, 100, 1010, 7050}, expected: []int{10}},
		{n: 0, input: []int{10, 15, 20, 100, 1010, 7050}, expected: []int{}},
		{n: -3, input: []int{10, 15, 20, 100, 1010, 7050}, expected: []int{}},
	}

	for test_number, test := range tests {
		if taken := Take(test.n, test.input); !ArrayEqual(taken, test.expected) {
			msg := strings.Join([]string{
				fmt.Sprintf("Test #%d failed on Take(%d, %v)\n", test_number+1, test.n, test.input),
				fmt.Sprintf("\tReturned: %v\n", taken),
				fmt.Sprintf("\tExpected: %v\n", test.expected),
			}, "")
			t.Error(msg)
		}
	}
}
