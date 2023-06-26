package models

import c "aoc/common"

type Point = c.Pair[int, int]

func MakePoint(height int, position int) Point {
	return c.MakePair(height, position)
}

func HeightOf(point Point) int {
	return point.First
}

func PositionOf(point Point) int {
	return point.Second
}
