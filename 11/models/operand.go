package models

type Operand interface {
	String() string
	eval(int) int
}
