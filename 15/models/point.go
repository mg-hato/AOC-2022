package models

import (
	f "aoc/functional"
	"fmt"
)

type Point = f.Pair[int, int]

func Abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func FormatPoint(point Point) string {
	return fmt.Sprintf("(%d, %d)", point.First, point.Second)
}

func MakePoint(x, y int) Point {
	return Point{First: x, Second: y}
}

// Manhattan distance
func Distance(pointX, pointY Point) int {
	return Abs(pointX.First-pointY.First) + Abs(pointX.Second-pointY.Second)
}
