package solver

import (
	"aoc/day05/models"
	f "aoc/functional"
)

type crate_mover_9000 struct{}

func CrateMover9000() CrateMover {
	return &crate_mover_9000{}
}

func (crate_mover_9000) String() string {
	return "CrateMover-9000"
}

func (crate_mover_9000) ExecuteMove(stacks []models.Containers, move models.Move) bool {
	source, destination := move.Source-1, move.Destination-1
	if len(stacks[source]) < move.Quantity {
		return false
	}
	stacks[destination] += string(f.Reverse([]rune(stacks[source][len(stacks[source])-move.Quantity:])))
	stacks[source] = stacks[source][:len(stacks[source])-move.Quantity]
	return true
}
