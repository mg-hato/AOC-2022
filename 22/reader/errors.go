package reader

import "fmt"

func reading_error_prefix(line_number int) string {
	return fmt.Sprintf(`error while reading line #%d`, line_number)
}

func validation_error_prefix() string {
	return "error while performing final validation"
}

func could_not_interpret_line_reading_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: could not interpret the line "%s"`,
		reading_error_prefix(line_number),
		line,
	)
}

func map_does_not_contain_walkable_paths_error() error {
	return fmt.Errorf(
		`%s: no position at the map can be walked on (i.e. no dots)`,
		validation_error_prefix(),
	)
}

func discountinuity_in_map_detected(location string) error {
	return fmt.Errorf(
		`%s: discountinuity detected in %s`,
		validation_error_prefix(), location,
	)
}
