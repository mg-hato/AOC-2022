package common_test

import (
	"aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestCommon_Sum(t *testing.T) {
	ts.AssertEqual(t, common.Sum([]int{1, 2, 34, 56, 10}), 103)
	ts.AssertEqual(t, common.Sum([]float64{1.5, 2.25, 3.25, 7.5}), 14.5)
}

func TestCommon_Abs(t *testing.T) {
	ts.AssertEqual(t, common.Abs(-10.5), 10.5)
	ts.AssertEqual(t, common.Abs(10), 10)
}
