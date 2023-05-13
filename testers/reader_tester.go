package testers

import (
	"testing"
)

type reader_tester[T any] struct {
	t *testing.T

	reader_func func(string) (T, error) // Reader function

	eq_func func(T, T) bool // Equality function for data read

	test_cases []reader_tester_test_case[T]
}

// Reader Tester constructor function
func ReaderTester[T any](t *testing.T, reader_func func(string) (T, error)) *reader_tester[T] {
	rt := reader_tester[T]{
		t: t,

		reader_func: reader_func,
		test_cases:  make([]reader_tester_test_case[T], 0),
	}
	return &rt
}

// Reader Tester constructor function for comparable data types that reader returns
func ReaderTesterForComparableData[T comparable](t *testing.T, reader_func func(string) (T, error)) *reader_tester[T] {
	return ReaderTester(t, reader_func).
		ProvideEqualityFunction(func(lhs, rhs T) bool { return lhs == rhs })
}

// Provide an equality function for data that reader returns
func (rt *reader_tester[T]) ProvideEqualityFunction(equality_func func(T, T) bool) *reader_tester[T] {
	rt.eq_func = equality_func
	return rt
}

// Provide a test case for reader tester i.e. an input filename and expected outcome of reading
func (rt *reader_tester[T]) AddTestCase(filename string, outcome expected_outcome[T]) *reader_tester[T] {
	test_case := reader_tester_test_case[T]{outcome, filename}
	test_case.addPrefix()
	rt.test_cases = append(rt.test_cases, test_case)
	return rt
}

func (rt *reader_tester[T]) RunReaderTests() *reader_tester[T] {
	if len(rt.test_cases) == 0 {
		rt.t.Error("There were no test cases provided; nothing to test")
		return rt
	} else if rt.eq_func == nil {
		rt.t.Error("No equality function was provided; tests cannot be run")
		return rt
	}
	for _, test_case := range rt.test_cases {
		data, err := rt.reader_func(test_case.filename)
		test_case.expected_outcome.testOutcome(rt.t, rt.eq_func, data, err)
	}
	return rt
}
