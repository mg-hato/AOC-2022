package models

type OldValue struct{}

func Old() Operand {
	return OldValue{}
}

func (OldValue) String() string {
	return "OLD"
}

func (OldValue) eval(old int) int {
	return old
}
