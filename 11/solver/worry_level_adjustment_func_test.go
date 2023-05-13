package solver

import (
	"aoc/testers"
	"testing"
)

func TestD11_LowestCommonMultiple(t *testing.T) {
	testers.AssertEqual(t, lcm(10, 15), 30)
	testers.AssertEqual(t, lcm(101, 7), 707)
	testers.AssertEqual(t, lcm(1, 6), 6)
	testers.AssertEqual(t, lcm(124, 400), 12_400)
}
