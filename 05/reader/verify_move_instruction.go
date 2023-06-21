package reader

import (
	c "aoc/common"
	"aoc/d05/models"
)

// Verify that move instruction is valid
func verify_move_instruction(move models.Move, number_of_stacks, line_number int) error {

	// 1. Verify that source and destination are in valid range
	is_in_range := c.InInclusiveRange(1, number_of_stacks)
	if !is_in_range(move.Source) || !is_in_range(move.Destination) {
		return source_or_destination_out_of_bounds_error(number_of_stacks, line_number)
	}

	// 2. Verify that source and destination are different stack IDs
	if move.Source == move.Destination {
		return same_source_and_destination_error(line_number)
	}

	return nil
}
