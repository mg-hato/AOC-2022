package solver

import (
	f "aoc/functional"
)

type WorryLevelAdjustmentStrategy int

const (
	DivBy3 WorryLevelAdjustmentStrategy = iota
	NoAdjustment
)

func getAdjustmentFunction(
	strategy WorryLevelAdjustmentStrategy,
	divisibility_test_numbers []int,
) func(int) int {
	// Modulo to prevent worry level going into overflow
	// This modulo maintains the outcomes of divisibility tests
	modulo := f.Foldl(lcm, divisibility_test_numbers, 1)

	switch strategy {
	case DivBy3:
		// If there is extra division by 3 happening after each analysis operation
		// then the modulo needs to be updated to maintain the outcomes of divisibility tests
		// accounting for the integer division by 3

		// Proof: say the actual worry level is W
		// Then it can be written as W = 3 * L * n + 3 * m + k where
		// - L is LCM of all divisibility test numbers
		// - n is a non-negative integer
		// - m is 0, 1, ... or L-1
		// - k is 0, 1 or 2
		// Let p be any of the divisibility numbers
		// N.B. that L % p = 0 by the merit of L being LCM of all divisibility numbers
		// Then W / 3 = L * n + m === m (mod p)
		// And (W % (3*L)) / 3 = (3 * m + k) / 3 = m
		// Hence the divisibility outcome is maintained for p
		// modulo *= 3

		return func(worry_level int) int { return (worry_level % (3 * modulo)) / 3 }
	default:
		return func(worry_level int) int { return worry_level % modulo }
	}
}

// Helper function: Least common multiplier
func lcm(i, j int) int {
	return i * j / gcd(i, j)
}

// Helper function: Greatest common divisor
func gcd(i, j int) int {
	if i == 0 || j == 0 {
		return 0
	}
	for j != 0 {
		if i < j {
			temp := i
			i = j
			j = temp
		} else {
			i %= j
		}
	}
	return i
}
