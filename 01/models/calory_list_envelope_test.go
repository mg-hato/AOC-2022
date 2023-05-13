package models

import (
	f "aoc/functional"
	ts "aoc/testers"
	"testing"
)

func TestD01_CaloryListEnvelope(t *testing.T) {
	envelope := CreateCaloryListEnvelope([][]int{{1, 2, 3}, {45}, {100, 5, 6}})

	// Make a change to the received data
	calory_list1 := envelope.Get()
	calory_list1[0][1] = 7

	// Ensure that data received through another Get call is unaffected by the change
	ts.AssertEqualWithEqFunc(t, envelope.Get(), [][]int{{1, 2, 3}, {45}, {100, 5, 6}}, f.ArrayEqualWith(f.ArrayEqual[int]))
}
