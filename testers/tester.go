package testers

import "testing"

type Tester[T any, R any] struct {
	tests []struct {
		input    T
		expected R
	}

	function      func(T) R
	function_name string

	equals func(R, R) bool
}

func DefaultTester[T any, R any](
	function func(T) R,
	function_name string,
) *Tester[T, R] {
	return &Tester[T, R]{
		function:      function,
		function_name: function_name,
		tests:         nil,
	}
}

func DefaultTesterForComparableTypeR[T any, R comparable](
	function func(T) R,
	function_name string,
) *Tester[T, R] {
	tester := DefaultTester(function, function_name)
	tester.equals = func(lhs, rhs R) bool { return lhs == rhs }
	return tester
}

func (tester *Tester[T, R]) ProvideEqualityFunction(equality func(R, R) bool) *Tester[T, R] {
	tester.equals = equality
	return tester
}

func (tester *Tester[T, R]) AddTest(input T, expected R) *Tester[T, R] {
	tester.tests = append(tester.tests, struct {
		input    T
		expected R
	}{input, expected})
	return tester
}

func (tester *Tester[T, R]) RunTests(t *testing.T) *Tester[T, R] {
	for test_number, test := range tester.tests {
		if tester.equals == nil {
			t.Error("Test failed. Equality function needs to be provided")
			return tester
		}
		if result := tester.function(test.input); !tester.equals(result, test.expected) {
			t.Errorf("Test #%d failed: %s(%v)", test_number, tester.function_name, test.input)
			t.Error("Actual result and expected result do not match")
			t.Errorf("Actual: %v", result)
			t.Errorf("Expected: %v", test.expected)
		}
	}
	return tester
}
