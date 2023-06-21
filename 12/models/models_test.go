package models

import (
	c "aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestD12_EnvelopeTest(t *testing.T) {
	data := func() Terrain { return []string{"Sabcdef", "mlkjihg", "nopqrsE"} }
	envelope := TerrainEnvelope(data()...)

	terrain := envelope.Get()
	terrain[0] = "xbw"

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data(), c.ArrayEqual[string])
}

func TestD12_GetNeighbours(t *testing.T) {
	ts.AssertEqualWithEqFunc(
		t,
		GetNeighbours(Position{First: 10, Second: 20}),
		[]Position{
			{First: 11, Second: 20},
			{First: 9, Second: 20},
			{First: 10, Second: 21},
			{First: 10, Second: 19},
		},
		c.ArrayEqualInAnyOrder[Position],
	)
}

func TestD12_EnumerateTerrain(t *testing.T) {
	ts.AssertEqualWithEqFunc(
		t,
		EnumerateTerrain([]string{"Sa", "ml", "sE"}),
		[][]Field{
			{{Position{First: 0, Second: 0}, 'S'}, {Position{First: 0, Second: 1}, 'a'}},
			{{Position{First: 1, Second: 0}, 'm'}, {Position{First: 1, Second: 1}, 'l'}},
			{{Position{First: 2, Second: 0}, 's'}, {Position{First: 2, Second: 1}, 'E'}},
		},
		c.ArrayEqualWith(c.ArrayEqual[Field]),
	)
}

func TestD12_GetHeight(t *testing.T) {
	ts.AssertEqual(t, Field{HeightCode: 'S'}.GetHeight(), 0)
	ts.AssertEqual(t, Field{HeightCode: 'a'}.GetHeight(), 0)
	ts.AssertEqual(t, Field{HeightCode: 'f'}.GetHeight(), 5)
	ts.AssertEqual(t, Field{HeightCode: 'z'}.GetHeight(), 25)
	ts.AssertEqual(t, Field{HeightCode: 'E'}.GetHeight(), 25)
}
