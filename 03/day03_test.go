package main

import (
	. "aoc/functional"
	"aoc/testers"
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
	tester := testers.DefaultReaderTester(ReadListOfContents, "ReadListOfContents")
	tester.ProvideEqualityFunctionForTypeT(func(lhs, rhs ListOfContents) bool {
		return ArrayEqual(lhs.rucksacks, rhs.rucksacks)
	})
	tester.AddGoodInputTests(
		loc("abcdef", "xyzw", "okok", "ab", "cd", "ef"),
		loc("aaaa", "pqrsaa", "ppqqppqq"),
		loc(),
	)
	tester.AddErrorInputTests(
		"There is an odd number of items on line 3 (number of items must be even)",
		"Line 3 has numbers, but only a-z and A-Z are allowed",
		"There are 4 rucksacks given, hence they cannot be split into groups of (exactly) three",
	)
	tester.RunBothGoodAndErrorInputTests(t)
}

func TestDay03_Solver(t *testing.T) {
	tester := testers.DefaultSolverTesterForComparableTypeR(
		func(list ListOfContents) (int, error) { return SumOfPriorities(list, FindRepeatedItems) },
		func(list ListOfContents) (int, error) { return SumOfPriorities(list, FindGroupBadges) },
		"SumOfPrioritiesOfRepeatedItems",
		"SumOfPrioritiesOfGroupBadges",
	)

	// Example input given in the problem statement
	tester.AddTest(
		loc(
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		), 157, 70,
	)

	// Custom input #1
	// Common items: x, P, j (24, 42, 10)
	// Badges: B (28)
	tester.AddTest(
		loc(
			"xyzBZXYx",     // xyz + B
			"PQRSTBpppsPs", // pqrst + B
			"ijkBjI",       // ijk + B
		), 76, 28,
	)

	// Custom input #2
	// Common items: A, f, H, k, O, q (27, 6, 34, 11, 41, 17)
	// Badges: W, z (49, 26)
	tester.AddTest(
		loc(
			"abcWbABBBAaB",   // abc + W
			"defdefDEfDEW",   // def + W
			"gHWGHi",         // ghi + W
			"jkzk",           // jk + z
			"mopnlzOOLLLLLL", // lmnop + z
			"rstquvqqqqqz",   // qrstuv + z
		), 136, 75,
	)

	tester.RunBothSolversTests(t)
}
