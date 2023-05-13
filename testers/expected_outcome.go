package testers

import "testing"

type expected_outcome[T any] interface {
	String() string

	// Tests whether the actual outcome matches expectation
	testOutcome(*testing.T, func(T, T) bool, T, error)

	// Sets prefix to be used in case of test failure
	setPrefix(string)
}
