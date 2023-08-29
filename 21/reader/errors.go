package reader

import "fmt"

func reading_error_prefix(line_number int) string {
	return fmt.Sprintf("error while reading on line #%d", line_number)
}

func final_validation_error_prefix() string {
	return "error while performing final validation"
}

func bad_line_reading_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: could not be interpret line "%s"`,
		reading_error_prefix(line_number),
		line,
	)
}

func repeated_id_reading_error(line_number int, id string, id_line_number int) error {
	return fmt.Errorf(
		`%s: monkey ID "%s" was already defined on line %d`,
		reading_error_prefix(line_number),
		id, id_line_number,
	)
}

func simple_job_is_not_a_number(line_number int) error {
	return fmt.Errorf(
		`%s: monkey job with no operations must have a number operand only`,
		reading_error_prefix(line_number),
	)
}

func operation_job_must_have_only_identifiers_reading_error(line_number int) error {
	return fmt.Errorf(
		`%s: monkey job with an operation must have identifier operands only`,
		reading_error_prefix(line_number),
	)
}

func unknown_operation_reading_error(line_number int, op_err error) error {
	return fmt.Errorf(
		`%s: %s`,
		reading_error_prefix(line_number),
		op_err.Error(),
	)
}

func monkey_id_missing_validation_error(id string) error {
	return fmt.Errorf(
		`%s: there is no monkey ID "%s"`,
		final_validation_error_prefix(), id,
	)
}

func self_reference_validation_error(id string, line_number int) error {
	return fmt.Errorf(
		`%s: calculation of monkey "%s" defined on line %d has a self reference`,
		final_validation_error_prefix(),
		id, line_number,
	)
}

func unknown_reference_validation_error(id string, line_number int, unknown_id string) error {
	return fmt.Errorf(
		`%s: calculation of monkey "%s" defined on line %d has an unknown reference "%s"`,
		final_validation_error_prefix(),
		id, line_number, unknown_id,
	)
}
