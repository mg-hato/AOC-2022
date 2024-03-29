package solver

import (
	c "aoc/common"
	m "aoc/d14/models"
)

type cave_system_with_floor struct {
	abyss_cave_system *cave_system_with_abyss
	floor_depth       int

	floor_left_end, floor_right_end int
}

func DefaultCaveSystemWithFloor(rock_structures []m.RockStructure) CaveSystemSimulator {
	cs := &cave_system_with_floor{
		abyss_cave_system: caveSystemWithAbyss(rock_structures, m.MakePoint(500, 0)),
	}

	extended_rock_structures := append(rock_structures, []m.Point{m.MakePoint(500, 0)})

	cs.floor_depth = 2 + c.Maximum(c.FlatMap(
		func(rock_structure m.RockStructure) []int { return c.Map(m.DepthOf, rock_structure) },
		extended_rock_structures,
	))

	column_values := c.FlatMap(
		func(rock_structure m.RockStructure) []int { return c.Map(m.ColumnOf, rock_structure) },
		extended_rock_structures,
	)

	cs.floor_left_end, cs.floor_right_end = c.Minimum(column_values)-2, c.Maximum(column_values)+2

	for column := cs.floor_left_end; column <= cs.floor_right_end; column++ {
		cs.abyss_cave_system.pillars[column] = append(cs.abyss_cave_system.pillars[column], m.MakePillar(cs.floor_depth))
	}
	return cs
}

func (cs *cave_system_with_floor) dropSandUnit() bool {
	if cs.abyss_cave_system.isOccupied(cs.abyss_cave_system.sand_source) {
		return false
	} else if cs.abyss_cave_system.dropSandUnit() {
		return true
	} else {
		// Extend the "infinite" floor
		cs.floor_left_end--
		cs.abyss_cave_system.pillars[cs.floor_left_end] = []m.Pillar{m.MakePillar(cs.floor_depth)}

		cs.floor_right_end++
		cs.abyss_cave_system.pillars[cs.floor_right_end] = []m.Pillar{m.MakePillar(cs.floor_depth)}

		return cs.abyss_cave_system.dropSandUnit()
	}
}
