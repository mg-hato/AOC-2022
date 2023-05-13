package models

type Multiplication struct{}

func Mult() Operator {
	return Multiplication{}
}

func (Multiplication) String() string {
	return "*"
}

func (Multiplication) apply(lhs, rhs int) int {
	return lhs * rhs
}
