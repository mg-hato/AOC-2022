package common_test

import (
	"aoc/common"
	"aoc/testers"
	"testing"
)

func TestCommon_SetEquality(t *testing.T) {
	s1 := map[int]bool{
		1: true,
		2: false,
		3: true,
	}
	s2 := map[int]bool{
		1: true,
		3: true,
		7: false,
	}

	testers.Assert(t, common.SetEqual(s1, s2))

	s2[2] = false
	testers.Assert(t, common.SetEqual(s1, s2))

	s2[2] = true
	testers.Assert(t, !common.SetEqual(s1, s2))
}

func TestCommon_CreateSet(t *testing.T) {
	set := common.CreateSet(
		[]int{1, 2, 3, 4},
		func(i int) int { return i * 2 },
	)

	testers.AssertEqualWithEqFunc(
		t,
		set,
		map[int]bool{2: true, 4: true, 6: true, 8: true},
		common.SetEqual[int],
	)

	set = common.CreateSet(
		[]int{1, 2, 3, 4},
		func(i int) int { return i / 2 },
	)

	testers.AssertEqualWithEqFunc(
		t,
		set,
		map[int]bool{0: true, 1: true, 2: true},
		common.SetEqual[int],
	)
}
