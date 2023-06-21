package solver

import (
	c "aoc/common"
	m "aoc/d14/models"
	"sort"
)

type cave_system_with_abyss struct {
	pillars map[int][]m.Pillar

	sand_source m.Point
}

func DefaultCaveSystemWithAbyss(rock_structures []m.RockStructure) CaveSystemSimulator {
	return caveSystemWithAbyss(rock_structures, m.MakePoint(500, 0))
}

func caveSystemWithAbyss(rock_structures []m.RockStructure, sand_source m.Point) *cave_system_with_abyss {
	cs := cave_system_with_abyss{
		sand_source: sand_source,

		pillars: make(map[int][]m.Pillar),
	}

	// Group rocks into "rock-sand pillars" per column (First axis)
	cs.pillars = c.GroupBy(
		m.GetRockPoints(rock_structures),
		c.GetFirst[int, int],
		func(rock_position m.Point) m.Pillar {
			return m.MakePillar(rock_position.Second)
		},
	)

	// Sort pillars per column
	for _, pillars := range cs.pillars {
		sort.Slice(pillars, func(i, j int) bool {
			return pillars[i].GetTop() < pillars[j].GetTop()
		})
	}

	return &cs
}

// Given a point, it returns the pillars that are in the same column as the point
func (cs cave_system_with_abyss) getPillarColumn(point m.Point) []m.Pillar {
	return cs.pillars[m.ColumnOf(point)]
}

// Given a point X it returns a tuple (index N, pillars P) s.t.
//   - the pillars P belong to the same column as point X
//   - the index N satisfies that P[0...N).tops <= X.depth < P[N...].tops
func (cs cave_system_with_abyss) pointBinarySearch(point m.Point) (int, []m.Pillar) {
	pillars := cs.getPillarColumn(point)
	left, right := 0, len(pillars)
	for left != right {
		if mid := (left + right) / 2; pillars[mid].GetTop() <= m.DepthOf(point) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left, pillars
}

// Returns true if and only if the given point is occupied by either rock or sand block
func (cs cave_system_with_abyss) isOccupied(point m.Point) bool {
	n, pillars := cs.pointBinarySearch(point)
	return n != 0 && pillars[n-1].ContainsDepth(m.DepthOf(point))
}

// Given a point X from which an object is dropped into free-fall it returns a tuple (boolean flag A, point N) s.t.
//   - boolean flag A is true if and only if the object would fall into abyss when dropped from X
//   - point N denoting the position of the object after the fall assuming it does not fall into abyss
func (cs cave_system_with_abyss) getPointAfterFall(point m.Point) (bool, m.Point) {
	n, pillars := cs.pointBinarySearch(point)
	if n == len(pillars) {
		return true, point
	} else {
		return false, m.MakePoint(m.ColumnOf(point), pillars[n].GetTop()-1)
	}
}

// Simulates a drop of sand from the sand-source and returns true if and only if sand-block has been
// successfully added to the cave system i.e. it did not fall into abyss
func (cs *cave_system_with_abyss) dropSandUnit() bool {
	if cs.isOccupied(cs.sand_source) {
		return false
	}

	sand, falls_into_abyss := cs.sand_source, false
	for {
		falls_into_abyss, sand = cs.getPointAfterFall(sand)

		if falls_into_abyss {
			return false
		}

		if next_point_left := m.MakePoint(m.ColumnOf(sand)-1, m.DepthOf(sand)+1); !cs.isOccupied(next_point_left) {
			sand = next_point_left
		} else if next_point_right := m.MakePoint(m.ColumnOf(sand)+1, m.DepthOf(sand)+1); !cs.isOccupied(next_point_right) {
			sand = next_point_right
		} else {
			cs.addSand(sand)
			return true
		}
	}
}

func (cs *cave_system_with_abyss) addSand(sand m.Point) {
	n, pillars := cs.pointBinarySearch(sand)
	pillars[n].AddSandBlock()
}
