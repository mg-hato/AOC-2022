package functional

import "testing"

func TestZip(t *testing.T) {

	p := func(i int, s string) Pair[int, string] {
		return Pair[int, string]{First: i, Second: s}
	}
	tests := []struct {
		inputA   []int
		inputB   []string
		expected []Pair[int, string]
	}{
		{
			[]int{10, 5, 11},
			[]string{"A", "X"},
			[]Pair[int, string]{p(10, "A"), p(5, "X")},
		},
		{
			[]int{},
			[]string{"A", "B", "C"},
			[]Pair[int, string]{},
		},
		{
			[]int{},
			[]string{},
			[]Pair[int, string]{},
		},
		{
			[]int{5, 5, 5},
			[]string{"1", "2", "3", "4"},
			[]Pair[int, string]{p(5, "1"), p(5, "2"), p(5, "3")},
		},
	}

	for test_number, test := range tests {
		if result := Zip(test.inputA, test.inputB); !ArrayEqual(result, test.expected) {
			t.Errorf("Test #%d failed: Zip(%v, %v)\n", test_number+1, test.inputA, test.inputB)
			t.Errorf("\tReturned: %v", result)
			t.Errorf("\tExpected: %v", test.expected)
		}
	}
}
