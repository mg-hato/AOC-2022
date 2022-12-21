package functional

import "testing"

func TestSum_integers(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 3}, 6},
		{[]int{-10, 20, -10, 1000}, 1000},
	}

	for test_number, test := range tests {
		if result := Sum(test.input); result != test.expected {
			t.Errorf("Test #%d failed: Sum(%v)", test_number+1, test.input)
			t.Errorf("\tReturned: %d", result)
			t.Errorf("\tExpected: %d", test.expected)
		}
	}
}

func TestSum_floats(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 0},
		{[]float64{1, 2, 3}, 6},
		{[]float64{-10, 20, -10, 1000}, 1000},
		{[]float64{0.25, -10, 20, 0.25, -10, 0.25}, 0.75},
	}

	for test_number, test := range tests {
		if result := Sum(test.input); result != test.expected {
			t.Errorf("Test #%d failed: Sum(%v)", test_number+1, test.input)
			t.Errorf("\tReturned: %v", result)
			t.Errorf("\tExpected: %v", test.expected)
		}
	}
}
