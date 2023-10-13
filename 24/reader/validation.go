package reader

import (
	c "aoc/common"
	"fmt"
)

func validate_map_is_not_empty(valley [][]rune) error {
	if len(valley) == 0 {
		return fmt.Errorf("empty map error")
	}
	return nil
}

func validate_square_map(valley [][]rune) error {
	row_length := len(valley[0])
	if c.Any(func(row []rune) bool { return len(row) != row_length }, valley) {
		return fmt.Errorf("map not square shaped")
	}
	return nil
}

func validate_sides_and_exits(valley [][]rune) error {
	err := fmt.Errorf("walls are not properly placed error")

	// left and right edges must be walls
	left, right := 0, len(valley[0])-1
	if c.Any(func(row []rune) bool { return row[left] != '#' || row[right] != '#' }, valley) {
		return err
	}

	count_char := func(chosen rune) func([]rune) int {
		return func(row []rune) int {
			return c.Count(row, func(r rune) bool { return r == chosen })
		}
	}

	count_walls := count_char('#')
	count_exits := count_char('.')

	// expected number of walls: all but one, which is reserved for exit/entrance
	expected_walls := len(valley[0]) - 1

	first_row, last_row := valley[0], valley[len(valley)-1]

	if count_exits(first_row) != 1 || count_walls(first_row) != expected_walls ||
		count_exits(last_row) != 1 || count_walls(last_row) != expected_walls {
		return err
	}
	return nil
}
