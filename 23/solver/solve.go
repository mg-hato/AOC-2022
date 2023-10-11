package solver

import (
	c "aoc/common"
	m "aoc/d23/models"
	"fmt"
)

func run_iterations(iterations int, envelope c.Envelope[[][]m.SpotType]) solver_result {

	current_elf_positions := to_nice_map(envelope.Get())

	i := 0
	done := false
	for !done && i < iterations {
		next_elf_positions := iteration(current_elf_positions, i)
		done = c.SetEqual(current_elf_positions, next_elf_positions)
		current_elf_positions = next_elf_positions
		i++
	}
	return solver_result{
		elves_settled:        done,
		elf_positions_at_end: current_elf_positions,
		iterations_finished:  i,
	}
}

func CountFreeSpacesInEncapsulatingRegion(iterations int) func(c.Envelope[[][]m.SpotType]) (int, error) {
	return func(envelope c.Envelope[[][]m.SpotType]) (int, error) {
		return count_free_spaces(run_iterations(iterations, envelope).elf_positions_at_end), nil
	}
}

func FirstRoundWhenNoElfMoves(max_iterations int) func(envelope c.Envelope[[][]m.SpotType]) (int, error) {
	return func(envelope c.Envelope[[][]m.SpotType]) (int, error) {
		result := run_iterations(max_iterations, envelope)
		var err error = nil
		if !result.elves_settled {
			err = fmt.Errorf(
				"error while solving: maximum number of iterations (%d) has been reached but the elves have not settled",
				max_iterations,
			)
		}
		return result.iterations_finished, err
	}
}

func count_free_spaces(elf_positions map[m.Position]bool) int {
	occupied_rows := c.Map(m.Position.Row, c.GetKeys(elf_positions))
	row_diff := c.Maximum(occupied_rows) - c.Minimum(occupied_rows)

	occupied_columns := c.Map(m.Position.Column, c.GetKeys(elf_positions))
	column_diff := c.Maximum(occupied_columns) - c.Minimum(occupied_columns)
	return (row_diff+1)*(column_diff+1) - len(elf_positions)
}

func get_iteration_movement_schedule(current_positions map[m.Position]bool, iteration_number int) func(current_position m.Position) m.Position {

	adjacent_empty := func(pos m.Position) func(m.Direction) bool {
		return func(dir m.Direction) bool {
			return !current_positions[pos.Adjacent(dir)]
		}
	}

	proposal_func := func(empty_dirs ...m.Direction) func(proposed m.Direction) func(pos m.Position) (bool, m.Position) {
		return func(proposed m.Direction) func(m.Position) (bool, m.Position) {
			return func(pos m.Position) (bool, m.Position) {
				return c.All(adjacent_empty(pos), empty_dirs), pos.Adjacent(proposed)
			}
		}
	}

	proposal_order := []func(m.Position) (bool, m.Position){
		proposal_func(m.North, m.NorthEast, m.NorthWest)(m.North),
		proposal_func(m.South, m.SouthEast, m.SouthWest)(m.South),
		proposal_func(m.West, m.NorthWest, m.SouthWest)(m.West),
		proposal_func(m.East, m.NorthEast, m.SouthEast)(m.East),
	}
	start := iteration_number % len(proposal_order)
	proposal_order = c.Flatten([][]func(m.Position) (bool, m.Position){proposal_order[start:], proposal_order[:start]})

	return func(current_position m.Position) m.Position {

		if c.All(adjacent_empty(current_position), m.GetAllDirections()) {
			return current_position
		}

		for _, proposal := range proposal_order {
			if satisfied, next_position := proposal(current_position); satisfied {
				return next_position
			}

		}
		return current_position
	}
}

func iteration(elf_positions map[m.Position]bool, iteration_number int) map[m.Position]bool {
	next := get_iteration_movement_schedule(elf_positions, iteration_number)

	grouped_by_next := c.GroupBy(c.GetKeys(elf_positions), next, c.Identity[m.Position])
	new_elf_positions := make(map[m.Position]bool)
	for next_pos, prev_positions := range grouped_by_next {
		if len(prev_positions) == 1 {
			new_elf_positions[next_pos] = true
		} else {
			c.ForEach(func(prev_pos m.Position) { new_elf_positions[prev_pos] = true }, prev_positions)
		}
	}
	return new_elf_positions
}

func to_nice_map(grove [][]m.SpotType) map[m.Position]bool {
	nm := map[m.Position]bool{}
	for i, row := range grove {
		for j, st := range row {
			if st == m.Elf {
				nm[m.MakePosition(i, j)] = true
			}
		}
	}
	return nm
}
