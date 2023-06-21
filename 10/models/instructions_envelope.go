package models

import (
	c "aoc/common"
)

type instructions_envelope struct {
	instructions []Instruction
}

func (envelope instructions_envelope) Get() []Instruction {
	return c.Map(c.Identity[Instruction], envelope.instructions)
}

func InstructionsEnvelope(instructions []Instruction) c.Envelope[[]Instruction] {
	return instructions_envelope{instructions}
}

func InstructionsEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]Instruction]) bool {
	return c.ArrayEqualWith(InstructionEqualityFunction)(lhs.Get(), rhs.Get())
}
