package reader

import "fmt"

func bad_line_reader_error(line_number int, line string) error {
	return fmt.Errorf(
		`error while reading line #%d: could not interpret the line "%s"`,
		line_number, line,
	)
}

func invalid_blueprint_ids_validation_error(n int) error {
	return fmt.Errorf(
		`%s: blueprint ids are supposed to be numbers from [1, %d] in ascending order`,
		"error while performing reader final validation",
		n,
	)
}
