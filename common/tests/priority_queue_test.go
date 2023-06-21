package common_test

import (
	c "aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestCommon_PriorityQueue(t *testing.T) {
	type record struct {
		id     string
		done   bool
		payoff int
	}

	// Use priority queue with priority function that
	// prioritises:
	// 1. records that are not marked as done
	// 2. if both records are not done, priority goes to the one with bigger payoff
	pq := c.PriorityQueue(func(lhs, rhs record) bool {
		if lhs.done {
			return false
		} else if rhs.done {
			return true
		}
		return lhs.payoff > rhs.payoff
	})

	pq.MultiAdd(
		record{id: "TOP", payoff: 1},
		record{id: "T", payoff: 6},
		record{id: "Z", payoff: 100},
		record{id: "WW", payoff: 21},
		record{id: "IJK", payoff: 41},
	)

	ts.AssertEqual(t, pq.Size(), 5)

	expected_order := []c.Pair[string, int]{
		c.MakePair("Z", 100),
		c.MakePair("IJK", 41),
		c.MakePair("WW", 21),
		c.MakePair("T", 6),
		c.MakePair("TOP", 1),
	}
	for i := 0; i < len(expected_order); i++ {
		R, err := pq.Pop()
		ts.AssertNoError(t, err)
		ts.AssertEqual(t, R.id, expected_order[i].First)
		ts.AssertEqual(t, R.payoff, expected_order[i].Second)
		ts.AssertEqual(t, R.done, false)
		R.done = true
		pq.Add(R)
	}

	R, err := pq.Pop()
	ts.AssertNoError(t, err)
	ts.AssertEqual(t, R.done, true)

	index := c.IndexOf(expected_order, func(pair c.Pair[string, int]) bool {
		return pair.First == R.id
	})
	ts.AssertEqual(t, R.payoff, expected_order[index].Second)
}
