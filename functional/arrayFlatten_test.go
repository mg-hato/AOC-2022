package functional

import "testing"

func TestFlatten(t *testing.T) {
	tests := []struct {
		input          [][]int
		expectedOutput []int
	}{
		{[][]int{{1, 2, 3}, {4, 5}, {6}}, []int{1, 2, 3, 4, 5, 6}},
		{[][]int{{}, {}, {}}, []int{}},
		{[][]int{{}, {10}, {}}, []int{10}},
	}

	for test_number, test := range tests {
		if result := Flatten(test.input); !ArrayEqual(result, test.expectedOutput) {
			t.Errorf("Test #%d failed. Flatten(%v) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}
