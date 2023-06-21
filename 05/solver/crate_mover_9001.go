package solver

import (
	"aoc/d05/models"
)

type crate_mover_9001 struct{}

func CrateMover9001() CrateMover {
	return &crate_mover_9001{}
}

func (crate_mover_9001) String() string {
	return "CrateMover-9001"
}

func (crate_mover_9001) ExecuteMove(stacks []models.Containers, move models.Move) bool {
	source, destination := move.Source-1, move.Destination-1
	if len(stacks[source]) < move.Quantity {
		return false
	}
	stacks[destination] += stacks[source][len(stacks[source])-move.Quantity:]
	stacks[source] = stacks[source][:len(stacks[source])-move.Quantity]
	return true
}
