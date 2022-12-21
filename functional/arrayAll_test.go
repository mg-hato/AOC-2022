package functional

import (
	"fmt"
	"strings"
	"testing"
)

func TestAll_func(t *testing.T) {

	isOdd := func(i int) bool {
		return i%2 != 0
	}

	isEven := func(i int) bool {
		return i%2 == 0
	}

	isNull := func(i int) bool {
		return i == 0
	}

	isNegative := func(i int) bool {
		return i < 0
	}

	isDivByN := func(n int) func(int) bool {
		return func(i int) bool {
			return i%n == 0
		}
	}

	tests := []struct {
		input           []int
		expected        bool
		predicate       func(int) bool
		funcDescription string
	}{
		{[]int{}, true, isNull, "isNull"},
		{[]int{1, 2, 3}, false, isNull, "isNull"},
		{[]int{1, 2, 0}, false, isNull, "isNull"},
		{[]int{1, 2, 0}, false, isNegative, "isNegative"},
		{[]int{-1}, true, isNegative, "isNegative"},
		{[]int{-1, -10, -7}, true, isNegative, "isNegative"},
		{[]int{1, 15, 7, 15}, false, isEven, "isEven"},
		{[]int{1, 15, 7, 15}, true, isOdd, "isOdd"},
		{[]int{0, 100, 67}, false, isOdd, "isOdd"},
		{[]int{5, 15, 75}, true, isDivByN(5), "isDivBy5"},
		{[]int{99, 0, 67}, false, isDivByN(99), "isDivBy99"},
	}

	for test_number, test := range tests {
		if result := All(test.predicate, test.input); result != test.expected {
			msg := strings.Join([]string{
				fmt.Sprintf("Test #%d failed: All(%s, %v)\n", test_number+1, test.funcDescription, test.input),
				fmt.Sprintf("\tReturned: %v\n", result),
				fmt.Sprintf("\tExpected: %v\n", test.expected),
			}, "")
			t.Error(msg)
		}
	}

}
