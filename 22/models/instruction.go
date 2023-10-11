package models

type Instruction interface {
	String() string
	GetMovement() int
	GetOrientation(current Orientation) Orientation
	Equals(Instruction) bool
}

func InstructionEqualityFunc(lhs, rhs Instruction) bool {
	return lhs.Equals(rhs)
}
