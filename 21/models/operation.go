package models

import (
	"fmt"
	"strings"
)

type Operation interface {
	Apply(int64, int64) (int64, error)
	ResolveRight(left, result int64) (int64, error)
	ResolveLeft(right, result int64) (int64, error)
	String() string
}

func PrintOperation(lhs string, op Operation, rhs string) string {
	return fmt.Sprintf("%s %s %s", lhs, op, rhs)
}

func TryParseOperation(op_str string) (Operation, error) {
	op_str = strings.TrimSpace(op_str)
	var op Operation
	var err error = nil
	switch op_str {
	case "+":
		op = Addition{}
	case "-":
		op = Subtraction{}
	case "*":
		op = Multiplication{}
	case "/":
		op = Division{}
	default:
		op = UndefinedOperation{}
		err = fmt.Errorf(`unrecognised operation: "%s"`, op_str)
	}
	return op, err
}
