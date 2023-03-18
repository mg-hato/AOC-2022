package main

import . "aoc/functional"

func AnalyseForestWith(getForestAnalyser func() ForestAnalyser) func([][]byte) (int, error) {
	return func(heights [][]byte) (int, error) {
		// "Trees" (i.e. heights with unique ids)
		trees := createForest(heights)

		// Create forest analyser
		forest_analyser := getForestAnalyser()

		ForEach(forest_analyser.AnalyseForestRow, trees)
		ForEach(forest_analyser.AnalyseForestRow, Map(Reverse[Tree], trees))
		ForEach(forest_analyser.AnalyseForestRow, transpose(trees))
		ForEach(forest_analyser.AnalyseForestRow, Map(Reverse[Tree], transpose(trees)))

		return forest_analyser.GetResult(), nil
	}
}

// Transpose 2D grid
func transpose(grid [][]Tree) [][]Tree {
	row_length := len(grid)
	column_length := len(grid[0])
	transposed := make([][]Tree, column_length)
	for i := 0; i < column_length; i++ {
		transposed[i] = make([]Tree, row_length)
	}

	for i := 0; i < row_length; i++ {
		for j := 0; j < column_length; j++ {
			transposed[j][i] = grid[i][j]
		}
	}
	return transposed
}

// Turn simple matrix of heights into matrix of "Tree" structs
func createForest(heights [][]byte) [][]Tree {
	var unique_id int
	return Map(func(bs []byte) []Tree {
		return Map(func(b byte) Tree {
			unique_id++
			return Tree{id: unique_id, height: b}
		},
			bs)
	}, heights)
}
