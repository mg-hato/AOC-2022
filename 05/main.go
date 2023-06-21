package main

import (
	"aoc/argshandle"
	"aoc/d05/reader"
	s "aoc/d05/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.MovingPlanReader),
		s.FollowMovingPlanWith(s.CrateMover9000()),
		s.FollowMovingPlanWith(s.CrateMover9001()),
	)
}
