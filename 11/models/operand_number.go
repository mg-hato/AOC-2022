package models

import "fmt"

type Number struct {
	number int
}

func Num(n int) Operand {
	return Number{n}
}

func (n Number) String() string {
	return fmt.Sprintf("NUMBER[%d]", n.number)
}

func (n Number) eval(int) int {
	return n.number
}
