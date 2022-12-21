package functional

import "testing"

func TestMinimum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{10, 5, 11}, 5},
		{[]int{5}, 5},
		{[]int{1, 2, -10}, -10},
	}

	lt := func(lhs, rhs int) bool { return lhs < rhs }

	for test_number, test := range tests {
		if result := Minimum(test.input, lt); result != test.expected {
			t.Errorf("Test #%d failed: Minimum(%v, <)\n", test_number+1, test.input)
			t.Errorf("\tReturned: %d", result)
			t.Errorf("\tExpected: %d", test.expected)
		}
	}
}
