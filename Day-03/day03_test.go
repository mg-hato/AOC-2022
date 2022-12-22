package main

import (
	. "aoc/functional"
	"fmt"
	"testing"
)

// shorth-hand for making list of contents
func loc(rucksacks ...string) ListOfContents {
	return ListOfContents{Map(
		func(s string) Rucksack {
			return Rucksack{s}
		}, rucksacks,
	)}
}

func TestDay03_Reader(t *testing.T) {
	tests := map[int]ListOfContents{
		1: loc("abcdef", "xyzw", "okok", "ab", "cd", "ef"),
		2: loc("aaaa", "pqrsaa", "ppqqppqq"),
		3: loc(),
	}

	for k, expected := range tests {
		filename := fmt.Sprintf("./test/input.%d", k)
		if result, _ := ReadListOfContents(filename); !result.equals(expected) {
			t.Errorf("Test #%d failed: ReadListOfContents(\"%s\")", k, filename)
			t.Errorf("Returned: %v", result)
			t.Errorf("Expected: %v", expected)
		}
	}
}

func (lhs ListOfContents) equals(rhs ListOfContents) bool {
	return ArrayEqual(lhs.rucksacks, rhs.rucksacks)
}

func TestDay03_ReaderErrors(t *testing.T) {
	bad_inputs := map[int]string{
		1: "There is an odd number of items on line 3 (number of items must be even)",
		2: "Line 3 has numbers, but only a-z and A-Z are allowed",
		3: "There are 4 rucksacks given, hence they cannot be split into groups of (exactly) three",
	}

	for number, reason := range bad_inputs {
		filename := fmt.Sprintf("./test/bad-input.%d", number)
		_, err := ReadListOfContents(filename)
		if err == nil {
			t.Errorf("Bad inputs test #%d failed. No error was returned when reading the file: \"%s\"", number, filename)
			t.Errorf("The input was bad because: %s", reason)
		}
	}

}

func TestDay03_SolverPart1(t *testing.T) {

	// Recall: a-z is of priority 1-26 and A-Z of 27-52
	tests := []struct {
		input    ListOfContents
		expected int
	}{
		{loc(), 0},
		{
			// Example input given in the problem statement
			loc(
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			),
			157,
		},
		{
			loc(
				"creaTeCRATES", // create | CRATES; T => 20 + 26 = 46
				"thirdLOVEr",   // third | LOVEr; r => 18
				"cookBOoK",     // cook | BOoK; o => 15
				"TeSTrest",     // TeST | rest; e => 5
			),
			46 + 18 + 15 + 5,
		},
		{loc("LONGWORDhaRdtest"), 44}, // LONGWORD | haRdtest; R => 18 + 26 = 44
	}

	for tn, test := range tests {
		if result := SumOfPriorities(&test.input, FindRepeatedItems); result != test.expected {
			t.Errorf("Test #%d failed: SumOfPriorities(%s)", tn+1, test.input)
			t.Errorf("Returned: %d", result)
			t.Errorf("Expected: %d", test.expected)
		}
	}
}

func TestDay03_SolverPart2(t *testing.T) {

	// Recall: a-z is of priority 1-26 and A-Z of 27-52
	tests := []struct {
		input    ListOfContents
		expected int
	}{
		{loc(), 0},
		{
			// Example input given in the problem statement
			loc(
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			),
			70,
		},
		{
			loc(
				"AbcBccaRbbAaab", "defDDRfFeeeE", "XxxyYYxyZzRyzyZyyZZX", // Group badge: R => 18 + 26 = 44
				"pAbcBccabbAaab", "defDDfFeepeE", "XxpxyYYxyZzyzyZyyZZX", // Group badge: p => 16
			),
			44 + 16,
		},
	}

	for tn, test := range tests {
		if result := SumOfPriorities(&test.input, FindGroupBadges); result != test.expected {
			t.Errorf("Test #%d failed: SumOfPriorities(%s)", tn+1, test.input)
			t.Errorf("Returned: %d", result)
			t.Errorf("Expected: %d", test.expected)
		}
	}
}
