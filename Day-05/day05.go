package main

import (
	"aoc/argshandle"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		ReadRearrangementPlan,
		func(plan RearrangementPlan) string { return FollowPlan(plan, CrateMover9000{}) },
		func(plan RearrangementPlan) string { return FollowPlan(plan, CrateMover9001{}) },
	)
}