package reader

import (
	"fmt"
)

const final_validation_error_prefix = "error while final validation was performed"

func invalid_line_error(line_number int, line string) error {
	return fmt.Errorf(
		`error while reading line #%d: could not interpret "%s"`,
		line_number, line,
	)
}

func no_terrain_was_given_error() error {
	return fmt.Errorf(
		"%s: not a single line was read that described terrain",
		final_validation_error_prefix,
	)
}

func terrain_map_is_not_rectangular_error() error {
	return fmt.Errorf(
		`%s: read terrain has different lengths across rows`,
		final_validation_error_prefix,
	)
}

func letter_expected_exactly_once_error(letter rune, actual_frequency int) error {

	letter_name := func() string {
		switch letter {
		case 'S':
			return "starting position"
		case 'E':
			return "finish position"
		default:
			return fmt.Sprintf("letter '%s'", string(letter))
		}
	}

	return fmt.Errorf(
		`%s: there should be exactly one %s, but there are %d instead`,
		final_validation_error_prefix,
		letter_name(),
		actual_frequency,
	)
}
