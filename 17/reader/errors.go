package reader

import (
	"errors"
	"fmt"
)

func bad_line_reader_error(line_number int, line string) error {
	return fmt.Errorf(
		`reader error encountered on line #%d: could not interpret line "%s"`,
		line_number, line,
	)
}

func jet_pattern_not_read_validation_error() error {
	return errors.New("final reader validation error: no jet pattern has been read")
}
