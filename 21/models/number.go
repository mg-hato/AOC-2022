package models

import "fmt"

type Number struct {
	number int64
}

func CreateNumber(n int64) Operand {
	return Number{n}
}

func (number Number) GetNumber() int64 {
	return number.number
}

func (number Number) String() string {
	return fmt.Sprintf("NUM[%d]", number.number)
}

func (thisNumber Number) Equal(other Operand) bool {
	otherNumber, ok := other.(Number)
	return ok && thisNumber.number == otherNumber.number
}

func (number Number) Resolve(func(string) (int64, error)) (int64, error) {
	return number.number, nil
}
