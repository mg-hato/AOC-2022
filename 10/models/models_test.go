package models

import (
	f "aoc/functional"
	ts "aoc/testers"
	"testing"
)

func TestD10_Envelope(t *testing.T) {
	data := func() []Instruction {
		return []Instruction{
			Addx(10),
			Noop(),
			Noop(),
			Addx(30),
			Noop(),
		}
	}
	envelope := InstructionsEnvelope(data())

	instructions := envelope.Get()
	instructions[0] = Noop()
	instructions[1] = Addx(5)

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data(), f.ArrayEqualWith(InstructionEqualityFunction))
}

func TestD10_Instructions(t *testing.T) {
	ts.AssertEqual(t, Addx(5).Execute(0), 5)
	ts.AssertEqual(t, Addx(10).Execute(5), 15)
	ts.AssertEqual(t, Noop().Execute(15), 15)
	ts.AssertEqual(t, Addx(-20).Execute(15), -5)
}
