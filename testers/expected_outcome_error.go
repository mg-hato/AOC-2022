package testers

import (
	f "aoc/functional"
	"fmt"
	"strings"
	"testing"
)

type expect_error[T any] struct {
	keywords []string
	prefix   string
}

// Public constructor for outcome that expects error that contains provided keywords
func ExpectError[T any](keywords ...string) expected_outcome[T] {
	return &expect_error[T]{
		keywords: keywords,
		prefix:   "",
	}
}

func (ee expect_error[T]) String() string {
	return fmt.Sprintf("ExpectError(keywords: %s)", format_string_array(ee.keywords))
}

func (ee *expect_error[T]) setPrefix(prefix string) {
	ee.prefix = prefix
}

// Test whether the actual outcome yields an error and whether such error contains all the expected keywords
func (ee expect_error[T]) testOutcome(t *testing.T, eq_func func(T, T) bool, result T, err error) {
	if err != nil {
		ee.assertThatErrorContainsKeywords(t, err)
	} else {
		t.Error(format_with_prefix(ee.prefix, "In ExpectError no error was returned"))
	}
}

func (ee expect_error[T]) assertThatErrorContainsKeywords(t *testing.T, err error) {
	err_msg_lowercase := strings.ToLower(err.Error())
	missing_keywords := f.Filter(
		func(keyword string) bool {
			return !strings.Contains(err_msg_lowercase, strings.ToLower(keyword))
		},
		ee.keywords,
	)
	if len(missing_keywords) > 0 {
		t.Error(format_with_prefix(ee.prefix,
			fmt.Sprintf(
				`In ExpectError test case the returned error message "%s" does not contain the following keywords: %s`,
				err.Error(),
				format_string_array(missing_keywords),
			),
		))
	}
}
