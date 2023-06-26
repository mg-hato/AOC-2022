package models

import c "aoc/common"

type Shape = []Point

func move_by(shape Shape, height_change, position_change int) Shape {
	return c.Map(func(p Point) Point {
		return MakePoint(HeightOf(p)+height_change, PositionOf(p)+position_change)
	}, shape)
}

func MoveLeft(shape Shape) Shape {
	return move_by(shape, 0, -1)
}

func MoveRight(shape Shape) Shape {
	return move_by(shape, 0, 1)
}

func MoveDown(shape Shape) Shape {
	return move_by(shape, -1, 0)
}
