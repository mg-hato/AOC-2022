package reader

import "fmt"

const moves_reading_error_prefix = "error reading move instructions"
const container_row_reading_error_prefix = "error reading container rows"

// Error when a line could not be reasonably interpreted while in reading container rows mode
func bad_line_reading_container_rows_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: could not interpret line #%d; line: "%s"`,
		container_row_reading_error_prefix,
		line_number,
		line,
	)
}

// Error when stack IDs are not unique integers from range [1, n] (where n is the number of stack IDs)
func invalid_stack_ids_error(number_of_stack_ids int) error {
	return fmt.Errorf(
		"%s: stack IDs should be unique integers from range [1, %d]",
		container_row_reading_error_prefix,
		number_of_stack_ids,
	)
}

// Error when container positions are not aligned with their respective stack IDs
func containers_not_aligned_with_stack_ids_error() error {
	return fmt.Errorf(
		`%s: the containers are not aligned with stack IDs`,
		container_row_reading_error_prefix,
	)
}

// Error when there are "floating" containers i.e. containers that have empty space underneath them
func floating_containers_error(floating_containers_stack_ids []int) error {
	return fmt.Errorf(
		`%s: floating containers detected on stacks: %v`,
		container_row_reading_error_prefix,
		floating_containers_stack_ids,
	)
}

// Error when a line could not be reasonably interpreted while in reading move instructions mode
func bad_line_reading_move_instructions_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: could not interpret line #%d; line: "%s"`,
		moves_reading_error_prefix,
		line_number,
		line,
	)
}

// Error when a move instruction is having source/destination stack ID out of expected range
func source_or_destination_out_of_bounds_error(number_of_stacks, line_number int) error {
	return fmt.Errorf(
		"%s: source/destination out of valid range [1, %d] on line #%d",
		moves_reading_error_prefix,
		number_of_stacks,
		line_number,
	)
}

// Error when a move instruction is having the same source and destination stack IDs
func same_source_and_destination_error(line_number int) error {
	return fmt.Errorf(
		"%s: same source and destination on line #%d",
		moves_reading_error_prefix,
		line_number,
	)
}
