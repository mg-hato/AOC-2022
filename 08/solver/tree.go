package solver

import (
	m "aoc/d08/models"
)

type tree struct {
	height byte
	pos    position
}

func enumerate_forest(forest m.Forest) [][]tree {
	enumerated_forest := make([][]tree, len(forest))
	for row_number := 0; row_number < len(forest); row_number++ {
		enumerated_forest[row_number] = make([]tree, len(forest[row_number]))
		for column_number := 0; column_number < len(forest[row_number]); column_number++ {
			enumerated_forest[row_number][column_number] = tree{
				height: forest[row_number][column_number],
				pos:    make_position(row_number, column_number),
			}
		}
	}
	return enumerated_forest
}
