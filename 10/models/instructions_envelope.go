package models

import (
	e "aoc/envelope"
	f "aoc/functional"
)

type instructions_envelope struct {
	instructions []Instruction
}

func (envelope instructions_envelope) Get() []Instruction {
	return f.Map(f.Identity[Instruction], envelope.instructions)
}

func InstructionsEnvelope(instructions []Instruction) e.Envelope[[]Instruction] {
	return instructions_envelope{instructions}
}

func InstructionsEnvelopeEqualityFunction(lhs, rhs e.Envelope[[]Instruction]) bool {
	return f.ArrayEqualWith(InstructionEqualityFunction)(lhs.Get(), rhs.Get())
}
