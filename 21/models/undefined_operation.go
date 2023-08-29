package models

import "fmt"

type UndefinedOperation struct {
	op_id string
}

func (undefined UndefinedOperation) Apply(int64, int64) (int64, error) {
	return 0, fmt.Errorf(`cannot apply undefined operation "%s"`, undefined.op_id)
}

func (UndefinedOperation) String() string {
	return "<?>"
}

func (undefined UndefinedOperation) ResolveRight(left, result int64) (int64, error) {
	return 0, could_not_resolve_right(undefined, left, result)
}

func (undefined UndefinedOperation) ResolveLeft(right, result int64) (int64, error) {
	return 0, could_not_resolve_left(undefined, right, result)
}
