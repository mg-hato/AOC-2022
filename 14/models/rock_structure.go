package models

import c "aoc/common"

type RockStructure = []Point

func GetRockPoints(rock_structures []RockStructure) []Point {
	min_and_max := func(lhs, rhs int) (int, int) {
		if lhs <= rhs {
			return lhs, rhs
		} else {
			return rhs, lhs
		}
	}

	connect_points := func(pointX, pointY Point) []Point {
		first_min, first_max := min_and_max(pointX.First, pointY.First)
		second_min, second_max := min_and_max(pointX.Second, pointY.Second)
		points := make([]Point, (first_max-first_min+1)*(second_max-second_min+1))
		index := 0
		for first := first_min; first <= first_max; first++ {
			for second := second_min; second <= second_max; second++ {
				points[index] = MakePoint(first, second)
				index++
			}
		}
		return points
	}

	rocks := map[Point]bool{}
	for _, rock_structure := range rock_structures {
		for i := 1; i < len(rock_structure); i++ {
			c.ForEach(
				func(p Point) { rocks[p] = true },
				connect_points(rock_structure[i-1], rock_structure[i]),
			)
		}
	}

	return c.GetKeys(rocks)
}
