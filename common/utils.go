package common

import "golang.org/x/exp/constraints"

func InRange(left, right int) func(int) bool {
	return func(n int) bool { return left <= n && n < right }
}

func InInclusiveRange(left, right int) func(int) bool {
	return func(n int) bool { return left <= n && n <= right }
}

// Constant function that returns `r`
func Const[T, R any](r R) func(T) R {
	return func(_ T) R { return r }
}

// Constant function that returns zero-value of type R
func ConstZero[T, R any](_ T) R {
	var r R
	return r
}

func GetZero[T any]() T {
	var zero_value T
	return zero_value
}

// Identity function
func Identity[T any](x T) T {
	return x
}

func Equal[T comparable](lhs, rhs T) bool {
	return lhs == rhs
}

func NotEqual[T comparable](lhs, rhs T) bool {
	return lhs != rhs
}

func LessThan[T constraints.Ordered](lhs, rhs T) bool {
	return lhs < rhs
}

func LessThanOrEqual[T constraints.Ordered](lhs, rhs T) bool {
	return lhs <= rhs
}

func GreaterThan[T constraints.Ordered](lhs, rhs T) bool {
	return lhs > rhs
}

func GreaterThanOrEqual[T constraints.Ordered](lhs, rhs T) bool {
	return lhs >= rhs
}

func Min[T constraints.Ordered](lhs, rhs T) T {
	if lhs <= rhs {
		return lhs
	} else {
		return rhs
	}
}

func Max[T constraints.Ordered](lhs, rhs T) T {
	if lhs >= rhs {
		return lhs
	} else {
		return rhs
	}
}

func Maximum[T constraints.Ordered](arr []T) T {
	return MaximumBy(arr, LessThan[T])
}

func Minimum[T constraints.Ordered](arr []T) T {
	return MinimumBy(arr, LessThan[T])
}
