package solver

import (
	c "aoc/common"
	m "aoc/d05/models"
	"fmt"
)

func not_enough_containers_error(move_number int, move m.Move) error {
	return fmt.Errorf(
		"error while executing move #%d %s: the source stack does not have enough containers",
		move_number, move,
	)
}

func empty_stacks_error(stacks []m.Containers) error {
	empty_stacks_ids := c.Map(c.GetFirst[int, m.Containers], c.Filter(
		func(pair c.Pair[int, m.Containers]) bool { return len(pair.Second) == 0 },
		c.EnumerateWithFirstIndex[string](1)(stacks),
	))
	return fmt.Errorf(
		"error following the moving plan: the following stacks have no containers %v",
		empty_stacks_ids,
	)
}
