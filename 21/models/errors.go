package models

import "fmt"

func could_not_resolve_right(op Operation, left, result int64) error {
	return fmt.Errorf(
		`error resolving right operand: could not solve for X in %d %s X = %d`,
		left, op.String(), result,
	)
}

func could_not_resolve_left(op Operation, right, result int64) error {
	return fmt.Errorf(
		`error resolving right operand: could not solve for X in X %s %d = %d`,
		op.String(), right, result,
	)
}
