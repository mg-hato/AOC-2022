package models

import c "aoc/common"

type fields_and_instructions_envelope struct {
	fields       []Field
	instructions []Instruction
}

func CreateFieldsAndInstructionsEnvelope(fields []Field, instructions []Instruction) c.Envelope[c.Pair[[]Field, []Instruction]] {
	return fields_and_instructions_envelope{fields, instructions}
}

func (fase fields_and_instructions_envelope) Get() c.Pair[[]Field, []Instruction] {
	return c.MakePair(
		c.ShallowCopy(fase.fields),
		c.ShallowCopy(fase.instructions),
	)
}

func FieldsAndInstructionsEnvelopeEqualityFunc(lhs, rhs c.Envelope[c.Pair[[]Field, []Instruction]]) bool {
	lhs_pair, rhs_pair := lhs.Get(), rhs.Get()
	return c.ArrayEqual(lhs_pair.First, rhs_pair.First) &&
		c.ArrayEqualWith(InstructionEqualityFunc)(lhs_pair.Second, rhs_pair.Second)
}
