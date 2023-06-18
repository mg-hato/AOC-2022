package models

import (
	f "aoc/functional"
	ts "aoc/testers"
	"testing"
)

func TestD14_EnvelopeTest(t *testing.T) {
	data := func() []RockStructure {
		return []RockStructure{
			{f.MakePair(200, 5), f.MakePair(200, 15), f.MakePair(250, 15)},
			{f.MakePair(100, 17), f.MakePair(111, 17), f.MakePair(123, 17), f.MakePair(123, 21)},
			{f.MakePair(910, 7), f.MakePair(910, 15)},
		}
	}

	envelope := RockStructureEnvelope(data())

	rs := envelope.Get()
	rs[0] = RockStructure{f.MakePair(7, 7), f.MakePair(7, 10)}
	rs[1][2].First = 119

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data(), f.ArrayEqualWith(f.ArrayEqual[f.Pair[int, int]]))
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
		return f.FlatMap(func(left int) []Point {
			return f.Map(func(right int) Point { return MakePoint(left, right) }, rhs)
		}, lhs)
	}
	expected_rocks := make(map[Point]bool)
	f.ForEach(func(p Point) { expected_rocks[p] = true }, f.Flatten([][]Point{
		cross_product(f.RangeInclusive(100, 100), f.RangeInclusive(20, 50)),

		cross_product(f.RangeInclusive(80, 110), f.RangeInclusive(50, 50)),
		cross_product(f.RangeInclusive(110, 110), f.RangeInclusive(45, 50)),
		cross_product(f.RangeInclusive(110, 110), f.RangeInclusive(45, 110)),

		cross_product(f.RangeInclusive(500, 507), f.RangeInclusive(31, 31)),
		cross_product(f.RangeInclusive(507, 507), f.RangeInclusive(21, 31)),
	}))
	rock_struct := []RockStructure{
		{MakePoint(100, 50), MakePoint(100, 20)},
		{MakePoint(80, 50), MakePoint(110, 50), MakePoint(110, 45), MakePoint(110, 110)},
		{MakePoint(500, 31), MakePoint(507, 31), MakePoint(507, 21)},
	}
	ts.AssertEqualWithEqFunc(
		t,
		GetRockPoints(rock_struct),
		f.GetKeys(expected_rocks),
		f.ArrayEqualInAnyOrder[Point],
	)
}
