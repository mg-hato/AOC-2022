package models

type Instruction interface {
	String() string

	Length() int     // length required in cycles to execute the instruction
	Execute(int) int // given current register's value returns new value

	equal(Instruction) bool
}

func InstructionEqualityFunction(lhs, rhs Instruction) bool {
	return lhs.equal(rhs)
}
