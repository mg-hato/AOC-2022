package models

import (
	c "aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestD02_RoundsEnvelope(t *testing.T) {
	env := CreateRoundsEnvelope([]Round{{A, X}, {B, Y}, {C, Z}})

	// Change contents of the first round
	rounds1 := env.Get()
	rounds1[0].Left = C

	// Check that subsequent Get() is unaffected by the change
	ts.AssertEqualWithEqFunc(t, env.Get(), []Round{{A, X}, {B, Y}, {C, Z}}, c.ArrayEqual[Round])
}
