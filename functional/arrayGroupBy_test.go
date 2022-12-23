package functional

import (
	"fmt"
	"testing"
)

func TestGroupBy(t *testing.T) {

	// Group numbers by the number of digits they have
	// E.g. numbers 3,6,7 would be in 1-digit numbers group and number 101, 999, 456 would be in 3-digit group

	// key-extractor function: digit counter function
	keyf := func(i int) int {
		return len(fmt.Sprint(i))
	}

	// value-extractor function: identity
	valf := Identity[int]

	tests := []struct {
		input    []int
		expected map[int][]int
	}{
		{
			[]int{1, 5, 20, 9, 1, 50, 100},
			map[int][]int{
				1: {1, 5, 9, 1},
				2: {20, 50},
				3: {100},
			},
		},
		{
			[]int{5005, 11, 5005, 11},
			map[int][]int{
				4: {5005, 5005},
				2: {11, 11},
			},
		},
	}

	for test_number, test := range tests {
		if result := GroupBy(test.input, keyf, valf); !areEqual(result, test.expected) {
			t.Errorf("Test #%d failed: GroupBy(%v, digit_counter, id_func)", test_number+1, test.input)
			t.Errorf("Returned: %v", result)
			t.Errorf("Expected: %v", test.expected)
		}
	}
}

func areEqual(lhs, rhs map[int][]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k := range lhs {
		if !ArrayEqual(lhs[k], rhs[k]) {
			return false
		}
	}
	return true
}
