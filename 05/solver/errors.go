package solver

import (
	m "aoc/day05/models"
	f "aoc/functional"
	"fmt"
)

func not_enough_containers_error(move_number int, move m.Move) error {
	return fmt.Errorf(
		"error while executing move #%d %s: the source stack does not have enough containers",
		move_number, move,
	)
}

func empty_stacks_error(stacks []m.Containers) error {
	empty_stacks_ids := f.Map(f.GetFirst[int, m.Containers], f.Filter(
		func(pair f.Pair[int, m.Containers]) bool { return len(pair.Second) == 0 },
		f.EnumerateWithFirstIndex(stacks, 1),
	))
	return fmt.Errorf(
		"error following the moving plan: the following stacks have no containers %v",
		empty_stacks_ids,
	)
}
