package common_test

import (
	"aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestCommon_Queue(t *testing.T) {
	queue := common.Queue[int](1, 2, 3)
	ts.AssertEqual(t, queue.Size(), 3)

	queue.Enqueue(4)
	for i := 1; i <= 4; i++ {
		popped, err := queue.Dequeue()
		ts.AssertNoError(t, err)
		ts.AssertEqual(t, popped, i)
	}
	_, err := queue.Dequeue()
	ts.AssertError(t, err)

	queue.Enqueue(4, 5, 6, 10, 11)
	ts.AssertEqualWithEqFunc(t, queue.GetAsArray(), []int{4, 5, 6, 10, 11}, common.ArrayEqual[int])
}
