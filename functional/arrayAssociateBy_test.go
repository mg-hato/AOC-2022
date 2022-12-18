package functional

import (
	"fmt"
	"testing"
)

func TestAssociateBy(t *testing.T) {

	// Maps everything to "x"
	x := struct {
		desc string
		f    func(int) string
	}{
		"_ -> x",
		func(i int) string { return "x" },
	}

	// Maps 1-26 => a-z, everything else to empty string
	letter := struct {
		desc string
		f    func(int) string
	}{
		"a-z",
		func(i int) string {
			if 1 <= i && i <= int('z'-'a'+1) {
				return fmt.Sprintf("%c", 'a'+i-1)
			} else {
				return ""
			}
		},
	}

	// Maps a number to its string format
	justString := struct {
		desc string
		f    func(int) string
	}{
		"string",
		func(i int) string { return fmt.Sprintf("%d", i) },
	}

	tests := []struct {
		input    []int
		expected map[string]int
		valf     struct {
			desc string
			f    func(int) string
		}
	}{
		{[]int{}, map[string]int{}, x},
		{[]int{5, 5, 5}, map[string]int{"x": 5}, x},
		{[]int{1, 5, -11}, map[string]int{"a": 1, "e": 5, "": -11}, letter},
		{[]int{1, 2, 3, 1, 2, 3}, map[string]int{"a": 1, "b": 2, "c": 3}, letter},
		{[]int{1, 2, 3, 1, 2, 3}, map[string]int{"1": 1, "2": 2, "3": 3}, justString},
		{[]int{15}, map[string]int{"15": 15}, justString},
	}

	for test_number, test := range tests {
		if result := AssociateBy(test.input, test.valf.f); !MapEqual(result, test.expected) {
			t.Errorf("Test #%d failed: AssociateBy(%v, %s)", test_number+1, test.input, test.valf.desc)
			t.Errorf("\tReturned: %v", result)
			t.Errorf("\tExpected: %v", test.expected)
		}
	}
}
