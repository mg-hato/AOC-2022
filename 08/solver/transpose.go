package solver

func transpose[T any](grid [][]T) [][]T {
	row_length := len(grid)
	column_length := len(grid[0])
	transposed := make([][]T, column_length)
	for i := 0; i < column_length; i++ {
		transposed[i] = make([]T, row_length)
	}

	for i := 0; i < row_length; i++ {
		for j := 0; j < column_length; j++ {
			transposed[j][i] = grid[i][j]
		}
	}
	return transposed
}
