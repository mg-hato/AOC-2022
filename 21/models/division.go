package models

import (
	"errors"
	"fmt"
)

type Division struct{}

func (Division) Apply(lhs, rhs int64) (int64, error) {
	if rhs == 0 {
		return 0, errors.New("division by zero")
	} else if rem := lhs % rhs; rem != 0 {
		return 0, fmt.Errorf("division has remainder %d", rem)
	} else {
		return lhs / rhs, nil
	}
}

func (Division) String() string {
	return "/"
}

func (d Division) ResolveRight(left, result int64) (int64, error) {
	// LHS / RHS = RESULT  ==>  RHS = LHS / RESULT
	if result == 0 || left == 0 || left%result != 0 {
		return 0, could_not_resolve_right(d, left, result)
	} else {
		return left / result, nil
	}
}

func (d Division) ResolveLeft(right, result int64) (int64, error) {
	// LHS / RHS = RESULT  ==> LHS = RESULT * RHS
	if right == 0 {
		return 0, could_not_resolve_left(d, right, result)
	}
	return right * result, nil
}
