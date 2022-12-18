package functional

import (
	"fmt"
	"strings"
	"testing"
)

func TestDrop_test(t *testing.T) {
	tests := []struct {
		n               int
		input, expected []int
	}{
		{n: 3, input: []int{1}, expected: []int{}},
		{n: 20, input: []int{1}, expected: []int{}},
		{n: 0, input: []int{1}, expected: []int{1}},
		{n: 3, input: []int{10, 15, 20, 100, 1010, 7050}, expected: []int{100, 1010, 7050}},
		{n: 5, input: []int{10, 15, 20, 100, 1010, 7050}, expected: []int{7050}},
		{n: -7, input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	}

	for test_number, test := range tests {
		if returned := Drop(test.n, test.input); !ArrayEqual(returned, test.expected) {
			msg := strings.Join([]string{
				fmt.Sprintf("Test #%d failed on Drop(%d, %v)\n", test_number+1, test.n, test.input),
				fmt.Sprintf("\tReturned: %v\n", returned),
				fmt.Sprintf("\tExpected: %v\n", test.expected),
			}, "")
			t.Error(msg)
		}
	}
}
