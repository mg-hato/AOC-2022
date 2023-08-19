package models

type Operand interface {
	String() string
	Equal(Operand) bool

	Resolve(func(string) (int64, error)) (int64, error)
}
