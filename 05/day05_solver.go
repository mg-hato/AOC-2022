package main

import (
	. "aoc/functional"
)

func FollowPlan(plan RearrangementPlan, crateMover CrateMover) string {
	stacks := CreateKeyValueMap(
		EnumerateWithFirstIndex(plan.stacks, 1),
		func(pair Pair[int, string]) int { return pair.First },
		func(pair Pair[int, string]) *Stack { return &Stack{pair.First, []rune(pair.Second)} },
	)

	for _, move := range plan.moves {
		crateMover.ApplyMove(move, stacks)
	}
	return string(Map(func(pair Pair[int, string]) rune { return stacks[pair.First].peek() }, EnumerateWithFirstIndex(plan.stacks, 1)))
}
