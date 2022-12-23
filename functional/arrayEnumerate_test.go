package functional

import "testing"

// short-hand for making int-string pairs
func p(i int, s string) Pair[int, string] {
	return Pair[int, string]{i, s}
}

func TestEnumerate(t *testing.T) {
	tests := []struct {
		input    []string
		expected []Pair[int, string]
		start    int
	}{
		{[]string{}, []Pair[int, string]{}, 10},
		{[]string{"A", "B", "C"}, []Pair[int, string]{p(1, "A"), p(2, "B"), p(3, "C")}, 1},
		{[]string{"aaaa", "yyy", "pqrst"}, []Pair[int, string]{p(1000, "aaaa"), p(1001, "yyy"), p(1002, "pqrst")}, 1000},
		{[]string{"a", "a", "a"}, []Pair[int, string]{p(-7, "a"), p(-6, "a"), p(-5, "a")}, -7},
	}

	for test_number, test := range tests {
		if result := EnumerateWithFirstIndex(test.input, test.start); !ArrayEqual(result, test.expected) {
			t.Errorf("Test #%d failed: EnumerateWithFirstIndex(%v, %d)", test_number+1, test.input, test.start)
			t.Errorf("Returned: %v", result)
			t.Errorf("Expected: %v", test.expected)
		}
	}

}

func TestEnumerate_default(t *testing.T) {
	tests := []struct {
		input    []string
		expected []Pair[int, string]
	}{
		{[]string{}, []Pair[int, string]{}},
		{[]string{"A", "B", "C"}, []Pair[int, string]{p(0, "A"), p(1, "B"), p(2, "C")}},
		{[]string{"a", "a", "a"}, []Pair[int, string]{p(0, "a"), p(1, "a"), p(2, "a")}},
		{[]string{}, []Pair[int, string]{}},
	}

	for test_number, test := range tests {
		if result := Enumerate(test.input); !ArrayEqual(result, test.expected) {
			t.Errorf("Test #%d failed: Enumerate(%v)", test_number+1, test.input)
			t.Errorf("Returned: %v", result)
			t.Errorf("Expected: %v", test.expected)
		}
	}

}
