package main

import (
	. "aoc/functional"
	"fmt"
	"testing"
)

// quick-way to initialise assignment pair
func ap(a, b, c, d int) AssignmentPair {
	return AssignmentPair{
		first:  SectionRange{left: a, right: b},
		second: SectionRange{left: c, right: d},
	}
}

func TestDay04_ReaderOk(t *testing.T) {
	tests := map[int][]AssignmentPair{
		1: {ap(1, 5, 10, 10), ap(711, 1000, 1, 1), ap(3, 4, 5, 5)},
		2: {ap(399, 401, 400, 400)},
		3: {}, // empty file, should be valid!
	}

	for k, expected := range tests {
		filename := fmt.Sprintf("./test/input.%d", k)
		if result, err := ReadListOfAssignmentPairs(filename); err != nil || !ArrayEqual(result, expected) {
			t.Errorf("Test #%d failed: ReadListOfAssignmentPairs(\"%s\")", k, filename)
			if err != nil {
				t.Errorf("An unexpected error occurred. Error details below:")
				t.Error(err.Error())
			} else {
				t.Errorf("Returned: %s", result)
				t.Errorf("Expected: %s", expected)
			}
		}
	}

}

func TestDay04_ReaderErrs(t *testing.T) {
	tests := map[int]string{
		1: "Line 3 defines too many ranges",
		2: "Line 4 defines invalid range",
		3: "Line 4 range defined poorly with not a number",
	}

	for k, reason := range tests {
		filename := fmt.Sprintf("./test/bad-input.%d", k)
		if _, err := ReadListOfAssignmentPairs(filename); err == nil {
			t.Errorf("Test #%d failed: ReadListOfAssignmentPairs(\"%s\")", k, filename)
			t.Errorf("No error was returned, but one was expected because: %s", reason)

		}
	}
}

func TestSolver(t *testing.T) {
	tests := []struct {
		input                []AssignmentPair
		expected1, expected2 int // Expected answer for task 1 & 2
	}{
		// Example from the problem statement
		{input: []AssignmentPair{
			ap(2, 4, 6, 8),
			ap(2, 3, 4, 5),
			ap(5, 7, 7, 9),
			ap(2, 8, 3, 7),
			ap(6, 6, 4, 6),
			ap(2, 6, 4, 8),
		}, expected1: 2, expected2: 4},
		{input: []AssignmentPair{
			ap(1, 1, 1, 2),         // F O
			ap(1, 1, 1, 2),         // F O
			ap(1, 5, 7, 9),         //
			ap(100, 200, 150, 250), // O
			ap(100, 200, 130, 201), // O
			ap(100, 200, 201, 201), //
			ap(51, 60, 55, 60),     // F O
			ap(2, 8, 3, 7),         // F O
			ap(6, 6, 4, 6),         // F O
			ap(2, 6, 4, 8),         // O
			ap(15, 167, 10, 200),   // F O
			ap(15, 167, 168, 200),  //
		}, expected1: 6, expected2: 9},
	}

	for tn, test := range tests {
		if result := CountAssignmentPairsThatSatisfy(OneFullyContainsTheOther, test.input); result != test.expected1 {
			t.Errorf("Test #%d failed: CountAssignmentPairsThatSatisfy(FullyContains, %s)", tn+1, test.input)
			t.Errorf("Returned: %d", result)
			t.Errorf("Expected: %d", test.expected1)
		}

		if result := CountAssignmentPairsThatSatisfy(SectionRangesOverlap, test.input); result != test.expected2 {
			t.Errorf("Test #%d failed: CountAssignmentPairsThatSatisfy(Overlaps, %s)", tn+1, test.input)
			t.Errorf("Returned: %d", result)
			t.Errorf("Expected: %d", test.expected2)
		}
	}

}
