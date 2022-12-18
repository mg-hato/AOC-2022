package functional

import (
	"fmt"
	"strings"
	"testing"
)

func TestAny_func(t *testing.T) {

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
		{[]int{}, false, isNull, "isNull"},
		{[]int{1, 2, 3}, false, isNull, "isNull"},
		{[]int{1, 2, 0}, true, isNull, "isNull"},
		{[]int{1, 2, 0}, false, isNegative, "isNegative"},
		{[]int{-1}, true, isNegative, "isNegative"},
		{[]int{1, 2, 10, 15}, false, isNegative, "isNegative"},
		{[]int{1, 15, 7, 15}, false, isEven, "isEven"},
		{[]int{0, 100, 67}, true, isOdd, "isOdd"},
		{[]int{1, 3, 6}, false, isDivByN(5), "isDivBy5"},
		{[]int{0, 100, 67}, true, isDivByN(99), "isDivBy99"},
	}

	for test_number, test := range tests {
		if result := Any(test.predicate, test.input); result != test.expected {
			msg := strings.Join([]string{
				fmt.Sprintf("Test #%d failed: Any(%s, %v)\n", test_number+1, test.funcDescription, test.input),
				fmt.Sprintf("\tReturned: %v\n", result),
				fmt.Sprintf("\tExpected: %v\n", test.expected),
			}, "")
			t.Error(msg)
		}
	}

}
