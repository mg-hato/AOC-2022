package reader

import (
	c "aoc/common"
	"fmt"
)

func reader_error_prefix(line_number int) string {
	return fmt.Sprintf("error while reading on line #%d", line_number)
}

func finald_validation_error_prefix() string {
	return "error while performing the final validation"
}

func bad_line_reading_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: could not interpret line "%s"`,
		reader_error_prefix(line_number),
		line,
	)
}

func rock_structures_are_not_horizontal_or_vertical_error(rock_formation_number, rock_line_number int) error {
	return fmt.Errorf(
		`%s: rock formation #%d rock line #%d is neither horizontal nor vertical`,
		finald_validation_error_prefix(),
		rock_formation_number,
		rock_line_number,
	)
}

func rock_formation_overlaps_sand_source_error(rock_formation_number int, sand_source c.Pair[int, int]) error {
	return fmt.Errorf(
		`%s: rock formation #%d overlaps the sand source (%d, %d)`,
		finald_validation_error_prefix(),
		rock_formation_number,
		sand_source.First,
		sand_source.Second,
	)
}
