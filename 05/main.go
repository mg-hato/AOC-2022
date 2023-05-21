package main

import (
	"aoc/argshandle"
	"aoc/day05/reader"
	s "aoc/day05/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.MovingPlanReader),
		s.FollowMovingPlanWith(s.CrateMover9000()),
		s.FollowMovingPlanWith(s.CrateMover9001()),
	)
}
