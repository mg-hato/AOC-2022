package solver

import (
	c "aoc/common"
	m "aoc/d14/models"
	ts "aoc/testers"
	"testing"
)

func rock_structures_test() []m.RockStructure {
	return []m.RockStructure{
		{m.MakePoint(498, 4), m.MakePoint(498, 6), m.MakePoint(496, 6)},
		{m.MakePoint(503, 4), m.MakePoint(502, 4), m.MakePoint(502, 9), m.MakePoint(494, 9)},
	}
}

func testCaveSystem() *cave_system_with_abyss {
	return caveSystemWithAbyss(rock_structures_test(), m.MakePoint(500, 0))
}

/*
			depth goes in this direction --->
			    0123456789
			494 .........#
			    .........#
			496 ......#..#
	            ......#..#
			498 ....###..#
	            .........#
			500 .........#
	            .........#
			502 ....######
	            ....#.....
			504 ..........
*/

func TestD14_CaveSystem(t *testing.T) {
	cs := testCaveSystem()

	expected_pillars := make(map[int][]m.Pillar)
	for column, depths := range map[int][]int{
		494: {9},
		495: {9},
		496: {6, 9},
		497: {6, 9},
		498: {4, 5, 6, 9},
		499: {9},
		500: {9},
		501: {9},
		502: {4, 5, 6, 7, 8, 9},
		503: {4},
	} {
		expected_pillars[column] = c.Map(m.MakePillar, depths)
	}

	dropSandAndAssert := func(drop_count int) {
		for i := 0; i < drop_count; i++ {
			ts.Assert(t, cs.dropSandUnit())
		}
		ts.AssertEqualWithEqFunc(
			t, cs.pillars, expected_pillars,
			c.MapEqualWith[int](c.ArrayEqual[m.Pillar]),
		)
	}

	setSandCount := func(pillar *m.Pillar, sand_count int) {
		*pillar = m.MakePillar(pillar.GetBase())
		for i := 0; i < sand_count; i++ {
			pillar.AddSandBlock()
		}
	}
	// Assert that the cave system is properly created
	dropSandAndAssert(0)

	// Assert the cave system handles properly 1 drop of sand
	setSandCount(&expected_pillars[500][0], 1)
	dropSandAndAssert(1)

	// Assert the cave system handles properly 2 drops of sand
	setSandCount(&expected_pillars[499][0], 1)
	dropSandAndAssert(1)

	// Assert the cave system handles properly 5 drops of sand
	setSandCount(&expected_pillars[500][0], 2)
	setSandCount(&expected_pillars[501][0], 1)
	setSandCount(&expected_pillars[499][0], 1)
	setSandCount(&expected_pillars[498][3], 1)
	dropSandAndAssert(3)

	// Assert the cave system handles properly 22 drops of sand
	setSandCount(&expected_pillars[500][0], 7)
	setSandCount(&expected_pillars[499][0], 6)
	setSandCount(&expected_pillars[501][0], 6)

	setSandCount(&expected_pillars[498][3], 2)
	setSandCount(&expected_pillars[497][1], 1)
	dropSandAndAssert(17)

	// Assert that the cave system handles properly 24 drops of sand
	setSandCount(&expected_pillars[497][0], 1)
	setSandCount(&expected_pillars[495][0], 1)
	dropSandAndAssert(2)

	// All subsequent sand should fall into the abyss
	for i := 0; i < 100; i++ {
		ts.Assert(t, !cs.dropSandUnit())
	}
}

func TestD14_BinarySearch(t *testing.T) {
	cs := testCaveSystem()

	i, p := cs.pointBinarySearch(m.MakePoint(500, 0))
	ts.AssertEqual(t, i, 0)
	ts.AssertEqualWithEqFunc(t, p, []m.Pillar{m.MakePillar(9)}, c.ArrayEqual[m.Pillar])

	i, p = cs.pointBinarySearch(m.MakePoint(496, 7))
	ts.AssertEqual(t, i, 1)
	ts.AssertEqualWithEqFunc(t, p, []m.Pillar{m.MakePillar(6), m.MakePillar(9)}, c.ArrayEqual[m.Pillar])

	i, p = cs.pointBinarySearch(m.MakePoint(503, 5))
	ts.AssertEqual(t, i, 1)
	ts.AssertEqualWithEqFunc(t, p, []m.Pillar{m.MakePillar(4)}, c.ArrayEqual[m.Pillar])
}

func TestD14_FallingTest(t *testing.T) {
	cs := testCaveSystem()

	ab, p := cs.getPointAfterFall(m.MakePoint(494, 5))
	ts.AssertEqual(t, ab, false)
	ts.AssertEqual(t, p, m.MakePoint(494, 8))

	ab, p = cs.getPointAfterFall(m.MakePoint(496, 5))
	ts.AssertEqual(t, ab, false)
	ts.AssertEqual(t, p, m.MakePoint(496, 5))

	ab, p = cs.getPointAfterFall(m.MakePoint(497, 7))
	ts.AssertEqual(t, ab, false)
	ts.AssertEqual(t, p, m.MakePoint(497, 8))

	ab, _ = cs.getPointAfterFall(m.MakePoint(504, 2))
	ts.AssertEqual(t, ab, true)

}

func TestD14_Occupied(t *testing.T) {
	cs := testCaveSystem()

	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(0, 0)), false)
	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(494, 7)), false)
	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(494, 9)), true)
	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(496, 6)), true)
	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(496, 7)), false)
	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(496, 8)), false)
	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(496, 9)), true)
	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(498, 4)), true)
	ts.AssertEqual(t, cs.isOccupied(m.MakePoint(498, 3)), false)
}

func TestD14_CaveSystemWithFloor(t *testing.T) {
	cave_system, _ := DefaultCaveSystemWithFloor(rock_structures_test()).(*cave_system_with_floor)

	getSandCount := func(p m.Pillar) int {
		return p.GetBase() - p.GetTop()
	}

	for i := 0; i < 93; i++ {
		ts.Assert(t, cave_system.dropSandUnit())
	}
	pillars := cave_system.abyss_cave_system.pillars
	ts.AssertEqual(t, getSandCount(pillars[500][0]), 9)
	ts.AssertEqual(t, getSandCount(pillars[501][0]), 8)
	ts.AssertEqual(t, getSandCount(pillars[499][0]), 8)

	ts.AssertEqual(t, getSandCount(pillars[498][0]), 2)
	ts.AssertEqual(t, getSandCount(pillars[498][3]), 2)

	ts.AssertEqual(t, getSandCount(pillars[490][0]), 1)
	ts.AssertEqual(t, getSandCount(pillars[510][0]), 1)
}
