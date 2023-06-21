package reader

import (
	c "aoc/common"
	"aoc/d08/models"
	"errors"
	"fmt"
)

func bad_line_error(line_number int, line string) error {
	return fmt.Errorf(
		`reader error on line #%d: could not interpret the line "%s"`,
		line_number, line,
	)
}

func empty_forest_error() error {
	return errors.New("reader error when finished reading: the read forest is empty")
}

func different_forest_row_lengths_error(forest models.Forest) error {
	return fmt.Errorf(
		"reader error when finished reading: forest rows have different lengths, concretely %v",
		c.Map(func(row []byte) int { return len(row) }, forest),
	)
}
