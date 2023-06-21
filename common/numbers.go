package common

import "golang.org/x/exp/constraints"

func Sum[T constraints.Integer | constraints.Float](arr []T) T {
	return Foldl(
		func(lhs, rhs T) T { return lhs + rhs },
		arr, 0,
	)
}

func Abs[T constraints.Signed | constraints.Float](val T) T {
	if val < 0 {
		return -val
	} else {
		return val
	}
}
