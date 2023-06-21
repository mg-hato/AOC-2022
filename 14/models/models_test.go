package models

import (
	c "aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestD14_EnvelopeTest(t *testing.T) {
	data := func() []RockStructure {
		return []RockStructure{
			{c.MakePair(200, 5), c.MakePair(200, 15), c.MakePair(250, 15)},
			{c.MakePair(100, 17), c.MakePair(111, 17), c.MakePair(123, 17), c.MakePair(123, 21)},
			{c.MakePair(910, 7), c.MakePair(910, 15)},
		}
	}

	envelope := RockStructureEnvelope(data())

	rs := envelope.Get()
	rs[0] = RockStructure{c.MakePair(7, 7), c.MakePair(7, 10)}
	rs[1][2].First = 119

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data(), c.ArrayEqualWith(c.ArrayEqual[c.Pair[int, int]]))
}

func TestD14_Pillar(t *testing.T) {
	p := MakePillar(98)
	ts.AssertEqual(t, p, Pillar{base: 98, sand_count: 0})
	ts.AssertEqual(t, p.ContainsDepth(98), true)
	ts.AssertEqual(t, p.ContainsDepth(97), false)
	ts.AssertEqual(t, p.ContainsDepth(99), false)

	// Add 3 sand blocks
	p.AddSandBlock()
	p.AddSandBlock()
	p.AddSandBlock()

	ts.AssertEqual(t, p.ContainsDepth(98), true)
	ts.AssertEqual(t, p.ContainsDepth(97), true)
	ts.AssertEqual(t, p.ContainsDepth(95), true)
	ts.AssertEqual(t, p.ContainsDepth(94), false)
	ts.AssertEqual(t, p.ContainsDepth(99), false)

}

func TestD14_GetRockPoints(t *testing.T) {
	cross_product := func(lhs, rhs []int) []Point {
		return c.FlatMap(func(left int) []Point {
			return c.Map(func(right int) Point { return MakePoint(left, right) }, rhs)
		}, lhs)
	}
	expected_rocks := make(map[Point]bool)
	c.ForEach(func(p Point) { expected_rocks[p] = true }, c.Flatten([][]Point{
		cross_product(c.RangeInclusive(100, 100), c.RangeInclusive(20, 50)),

		cross_product(c.RangeInclusive(80, 110), c.RangeInclusive(50, 50)),
		cross_product(c.RangeInclusive(110, 110), c.RangeInclusive(45, 50)),
		cross_product(c.RangeInclusive(110, 110), c.RangeInclusive(45, 110)),

		cross_product(c.RangeInclusive(500, 507), c.RangeInclusive(31, 31)),
		cross_product(c.RangeInclusive(507, 507), c.RangeInclusive(21, 31)),
	}))
	rock_struct := []RockStructure{
		{MakePoint(100, 50), MakePoint(100, 20)},
		{MakePoint(80, 50), MakePoint(110, 50), MakePoint(110, 45), MakePoint(110, 110)},
		{MakePoint(500, 31), MakePoint(507, 31), MakePoint(507, 21)},
	}
	ts.AssertEqualWithEqFunc(
		t,
		GetRockPoints(rock_struct),
		c.GetKeys(expected_rocks),
		c.ArrayEqualInAnyOrder[Point],
	)
}
