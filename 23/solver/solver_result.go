package solver

import "aoc/d23/models"

type solver_result struct {
	elves_settled        bool
	elf_positions_at_end map[models.Position]bool
	iterations_finished  int
}
