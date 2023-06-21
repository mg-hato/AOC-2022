package models

import c "aoc/common"

type Point = c.Pair[int, int]

func MakePoint(column, depth int) Point {
	return Point{
		First:  column,
		Second: depth,
	}
}

func ColumnOf(point Point) int {
	return point.First
}

func DepthOf(point Point) int {
	return point.Second
}
