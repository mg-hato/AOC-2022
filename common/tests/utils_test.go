package common_test

import (
	"aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestCommon_InRange(t *testing.T) {
	in_range := common.InRange(10, 20)
	ts.Assert(t, in_range(10))
	ts.Assert(t, in_range(19))
	ts.Assert(t, in_range(13))

	ts.Assert(t, !in_range(20))
	ts.Assert(t, !in_range(225))
	ts.Assert(t, !in_range(9))

	in_range = common.InRange(10, 10)
	for i := 0; i < 100; i++ {
		ts.Assert(t, !in_range(i))
	}
}
func TestCommon_InInclusiveRange(t *testing.T) {
	in_range := common.InInclusiveRange(10, 20)
	ts.Assert(t, in_range(10))
	ts.Assert(t, in_range(19))
	ts.Assert(t, in_range(13))
	ts.Assert(t, in_range(20))

	ts.Assert(t, !in_range(225))
	ts.Assert(t, !in_range(9))

	in_range = common.InInclusiveRange(10, 10)
	for i := 0; i < 100; i++ {
		if i != 10 {
			ts.Assert(t, !in_range(i))
		}
	}
	ts.Assert(t, in_range(10))
}

func TestCommon_Const(t *testing.T) {
	constf := common.Const[string](10)
	ts.AssertEqual(t, constf("aa"), 10)
	ts.AssertEqual(t, constf(""), 10)
	ts.AssertEqual(t, constf("DEF"), 10)

	constf = common.ConstZero[string, int]
	ts.AssertEqual(t, constf("aa"), 0)
	ts.AssertEqual(t, constf(""), 0)
	ts.AssertEqual(t, constf("DEF"), 0)
}

func TestCommon_GetZero(t *testing.T) {
	ts.AssertEqual(t, common.GetZero[string](), "")
	ts.AssertEqual(t, common.GetZero[int](), 0)
	ts.AssertEqual(t, common.GetZero[*[]int](), nil)
	ts.AssertEqual(t, common.GetZero[map[int]int]() == nil, true)
	ts.AssertEqual(t, common.GetZero[bool](), false)
}

func TestCommon_MinMax(t *testing.T) {
	ts.AssertEqual(t, common.Min(10, 20), 10)
	ts.AssertEqual(t, common.Min(10.75, 10.5), 10.5)
	ts.AssertEqual(t, common.Min("ZERO", "ALPHA"), "ALPHA")

	ts.AssertEqual(t, common.Max(10, 20), 20)
	ts.AssertEqual(t, common.Max(10.75, 10.5), 10.75)
	ts.AssertEqual(t, common.Max("ZERO", "ALPHA"), "ZERO")
}

func TestCommon_MinimumMaximum(t *testing.T) {
	numbers := []int{10, 4, 16, 200, 101, 5, 99}
	strs := []string{"tt", "aaa", "xyz", "def", "dddd", "z", "ok", "ko"}

	ts.AssertEqual(t, common.Minimum(numbers), 4)
	ts.AssertEqual(t, common.Maximum(numbers), 200)

	ts.AssertEqual(t, common.Minimum(strs), "aaa")
	ts.AssertEqual(t, common.Maximum(strs), "z")
}
