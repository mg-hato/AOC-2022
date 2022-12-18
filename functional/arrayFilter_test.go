package functional

import "testing"

func TestFilter_Odds(t *testing.T) {
	tests := []struct {
		input, expectedOutput []int
	}{
		{input: []int{2, 3, 4, 5, 6, 7, 8, 10}, expectedOutput: []int{3, 5, 7}},
		{input: []int{0, 2, 4, 100, 200}, expectedOutput: []int{}},
		{input: []int{}, expectedOutput: []int{}},
		{input: []int{77, 65, 101}, expectedOutput: []int{77, 65, 101}},
		{input: []int{3, 15, 1001, 1024, 1023}, expectedOutput: []int{3, 15, 1001, 1023}},
	}

	isOdd := func(i int) bool { return i%2 != 0 }

	for test_number, test := range tests {
		if result := Filter(isOdd, test.input); !ArrayEqual(result, test.expectedOutput) {
			t.Errorf("Test #%d failed. Filter(isOdd, %v) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}

func TestFilter_TwoCharacterStrings(t *testing.T) {
	tests := []struct {
		input, expectedOutput []string
	}{
		{input: []string{}, expectedOutput: []string{}},
		{input: []string{"Aa", "abc", "bb", "two"}, expectedOutput: []string{"Aa", "bb"}},
		{input: []string{"one", "dos", "quatro"}, expectedOutput: []string{}},
		{input: []string{"ok", "Hi", "!?", "so"}, expectedOutput: []string{"ok", "Hi", "!?", "so"}},
	}

	isTwoCharacter := func(s string) bool { return len(s) == 2 }

	for test_number, test := range tests {
		if result := Filter(isTwoCharacter, test.input); !ArrayEqual(result, test.expectedOutput) {
			t.Errorf("Test #%d failed. Filter(isTwoCharacter, %v) returned %v, but %v was expected",
				test_number+1, test.input,
				result, test.expectedOutput)
		}
	}
}
