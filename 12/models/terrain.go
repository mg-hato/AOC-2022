package models

type Terrain = []string

func EnumerateTerrain(terrain Terrain) [][]Field {
	enumerated_terrain := make([][]Field, len(terrain))
	for row := 0; row < len(terrain); row++ {
		enumerated_terrain[row] = make([]Field, len(terrain[row]))
		for col := 0; col < len(terrain[row]); col++ {
			enumerated_terrain[row][col] = Field{
				Position: Position{
					First:  row,
					Second: col,
				},
				HeightCode: rune(terrain[row][col]),
			}
		}
	}
	return enumerated_terrain
}
