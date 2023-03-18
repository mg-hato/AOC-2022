package testers

import (
	"testing"
)

type SolverTester[T any, R any] struct {
	tests []struct {
		input             T
		expected_result_1 R
		expected_result_2 R
	}

	solver_1_func_name string
	sovler_2_func_name string

	solver_1 func(T) (R, error)
	solver_2 func(T) (R, error)

	equals func(R, R) bool
}

// Default Solver tester with some predefined values.
// User is expected to define equality-function for result-type R.
func DefaultSolverTester[T any, R any](solver_1, solver_2 func(T) (R, error), solver_1_name, solver_2_name string) SolverTester[T, R] {
	return SolverTester[T, R]{
		tests: []struct {
			input             T
			expected_result_1 R
			expected_result_2 R
		}{},

		solver_1: solver_1,
		solver_2: solver_2,

		solver_1_func_name: solver_1_name,
		sovler_2_func_name: solver_2_name,
	}
}

func (tester *SolverTester[T, R]) ProvideEqualityFunctionForTypeR(equality_func func(R, R) bool) *SolverTester[T, R] {
	tester.equals = equality_func
	return tester
}

// Default Solver tester with some predefined values.
// The equality-function for result-type R comes out of the box.
func DefaultSolverTesterForComparableTypeR[T any, R comparable](solver_1, solver_2 func(T) (R, error), solver_1_name, solver_2_name string) SolverTester[T, R] {
	tester := DefaultSolverTester(solver_1, solver_2, solver_1_name, solver_2_name)
	tester.equals = func(lhs, rhs R) bool { return lhs == rhs }
	return tester
}

// Add test for the solver
func (tester *SolverTester[T, R]) AddTest(input T, expected_result_1, expected_result_2 R) *SolverTester[T, R] {
	tester.tests = append(tester.tests, struct {
		input             T
		expected_result_1 R
		expected_result_2 R
	}{
		input:             input,
		expected_result_1: expected_result_1,
		expected_result_2: expected_result_2,
	})
	return tester
}

// Retrieves the appropriate solver based on the `solver_id`
func (tester *SolverTester[T, R]) getSolver(solver_id int) func(T) (R, error) {
	if solver_id == 1 {
		return tester.solver_1
	} else {
		return tester.solver_2
	}
}

// Retrieve the appropriate solver function name based on the `solver_id`
func (tester *SolverTester[T, R]) getSolverFuncName(solver_id int) string {
	if solver_id == 1 {
		return tester.solver_1_func_name
	} else {
		return tester.sovler_2_func_name
	}
}

// Run solver test based on the solver_id
func (tester *SolverTester[T, R]) testSolver(t *testing.T, solver_id int) *SolverTester[T, R] {
	for i, test := range tester.tests {
		test_number := i + 1
		var expected_result R

		// Get the expected result
		if solver_id == 1 {
			expected_result = test.expected_result_1
		} else {
			expected_result = test.expected_result_2
		}

		// If the result and the expected result do not match: test fails
		result, e := tester.getSolver(solver_id)(test.input)
		if e != nil {
			t.Errorf("Test for solver %d #%d failed: %s(%v)", solver_id, test_number, tester.getSolverFuncName(solver_id), test.input)
			t.Errorf("Unexpected error occurred: %v", e)
		} else if !tester.equals(result, expected_result) {
			t.Errorf("Test for solver %d #%d failed: %s(%v)", solver_id, test_number, tester.getSolverFuncName(solver_id), test.input)
			t.Errorf("Returned: %v", result)
			t.Errorf("Expected: %v", expected_result)
		}
	}
	return tester
}

// Run tests for the first solver
func (tester *SolverTester[T, R]) RunFirstSolverTests(t *testing.T) *SolverTester[T, R] {
	return tester.testSolver(t, 1)
}

// Run tests for the second solver
func (tester *SolverTester[T, R]) RunSecondSolverTests(t *testing.T) *SolverTester[T, R] {
	return tester.testSolver(t, 2)
}

// Run tests for both solvers (solver 1 and 2)
func (tester *SolverTester[T, R]) RunBothSolversTests(t *testing.T) *SolverTester[T, R] {
	return tester.
		RunFirstSolverTests(t).
		RunSecondSolverTests(t)
}
