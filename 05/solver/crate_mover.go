package solver

import (
	"aoc/d05/models"
)

type CrateMover interface {
	String() string

	// Updates the container stacks according to the move passed.
	// It returns boolean indicating whether the move was executed successfully.
	ExecuteMove([]models.Containers, models.Move) bool
}
