package models

import "fmt"

func Addx(v int) Instruction {
	return addx{v}
}

type addx struct {
	new_value int
}

func (a addx) String() string {
	return fmt.Sprintf("ADDX %d", a.new_value)
}

func (a addx) Execute(register int) int {
	return register + a.new_value
}

func (addx) Length() int {
	return 2
}

func (a addx) equal(instruction Instruction) bool {
	if other_addx, ok := instruction.(addx); ok {
		return a.new_value == other_addx.new_value
	}
	return false
}
