package reader

import (
	"fmt"
)

func bad_line_reader_error(line_number int, line string) error {
	return fmt.Errorf(
		`error while reading line #%d: could not interpret line "%s"`,
		line_number, line,
	)
}

func invalid_zero_count_validation_error(actual_zero_count int) error {
	return fmt.Errorf(
		"error while performing reader final validation: the encrypted file must have exactly 1 zero, but it actually has %d",
		actual_zero_count,
	)
}
