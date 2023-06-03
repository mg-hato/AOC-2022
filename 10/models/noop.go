package models

func Noop() Instruction {
	return noop{}
}

type noop struct{}

func (noop) String() string {
	return "NOOP"
}

func (noop) Execute(register int) int {
	return register
}

func (noop) Length() int {
	return 1
}

func (noop) equal(instruction Instruction) bool {
	_, ok := instruction.(noop)
	return ok
}
