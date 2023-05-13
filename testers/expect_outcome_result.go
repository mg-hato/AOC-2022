package testers

import (
	"fmt"
	"testing"
)

type expect_result[T any] struct {
	expected_result T
	prefix          string
}

// Public constructor for outcome that expects certain result
func ExpectResult[T any](expected_result T) expected_outcome[T] {
	return &expect_result[T]{
		expected_result: expected_result,
		prefix:          "",
	}
}

func (er expect_result[T]) String() string {
	return fmt.Sprintf("ExpectResult(%v)", er.expected_result)
}

func (er *expect_result[T]) setPrefix(prefix string) {
	er.prefix = prefix
}

func (er expect_result[T]) testOutcome(t *testing.T, eq_func func(T, T) bool, result T, err error) {
	if err != nil {
		t.Error(format_with_prefix(er.prefix,
			fmt.Sprintf(
				`In ExpectResult test case an error was returned. Error message: "%s"`,
				err.Error(),
			),
		))
	} else if !eq_func(result, er.expected_result) {
		t.Error(format_with_prefix(er.prefix,
			fmt.Sprintf(
				`In ExpectResult test case the actual (%v) does not match the expected result (%v)`,
				result, er.expected_result,
			),
		))
	}
}
