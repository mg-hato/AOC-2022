package common_test

import (
	"aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestCommon_Stack(t *testing.T) {
	stack := common.MakeStack[int](1, 2, 3)
	ts.AssertEqual(t, stack.Size(), 3)

	stack.Push(4)
	for i := 0; i < 4; i++ {
		popped, err := stack.Pop()
		ts.AssertNoError(t, err)
		ts.AssertEqual(t, popped, 4-i)
	}
	_, err := stack.Pop()
	ts.AssertError(t, err)

	stack.Push(4, 5, 6, 10, 11)
	ts.AssertEqualWithEqFunc(t, stack.GetAsArray(), []int{11, 10, 6, 5, 4}, common.ArrayEqual[int])
}
