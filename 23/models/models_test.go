package models

import (
	ts "aoc/testers"
	"testing"
)

func TestD23_Adjacent(t *testing.T) {
	ts.AssertEqual(t, MakePosition(10, 10).Adjacent(North), MakePosition(9, 10))
	ts.AssertEqual(t, MakePosition(-1, 0).Adjacent(NorthWest), MakePosition(-2, -1))
	ts.AssertEqual(t, MakePosition(100, 100).Adjacent(East), MakePosition(100, 101))
	ts.AssertEqual(t, MakePosition(5, 5).Adjacent(South), MakePosition(6, 5))
	ts.AssertEqual(t, MakePosition(10, 10).Adjacent(SouthEast), MakePosition(11, 11))
	ts.AssertEqual(t, MakePosition(-30, 0).Adjacent(SouthWest), MakePosition(-29, -1))
}
