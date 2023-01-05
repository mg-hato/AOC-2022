package testers

import (
	"testing"
)

type integration_tester[T, R any] struct {
	t *testing.T

	reader_func func(string) (T, error)

	solvers []func(T) (R, error)

	test_cases []integration_tester_test_case[R]

	can_add_solvers bool

	equality_func func(R, R) bool
}

// Integration Tester constructor function
func IntegrationTester[T, R any](t *testing.T) *integration_tester[T, R] {
	return &integration_tester[T, R]{
		t: t,

		reader_func: nil,

		solvers: make([]func(T) (R, error), 0),

		test_cases: make([]integration_tester_test_case[R], 0),

		can_add_solvers: true,

		equality_func: nil,
	}
}

// Integration Tester constructor function for comparable result types
func IntegrationTesterForComparableResults[T any, R comparable](t *testing.T) *integration_tester[T, R] {
	return IntegrationTester[T, R](t).ProvideEqualityFunctionForResults(func(lhs, rhs R) bool { return lhs == rhs })
}

// Provide an equality function for results.
func (it *integration_tester[T, R]) ProvideEqualityFunctionForResults(eq_func func(R, R) bool) *integration_tester[T, R] {
	it.equality_func = eq_func
	return it
}

// Provide a reader function.
func (it *integration_tester[T, R]) ProvideReader(reader func(string) (T, error)) *integration_tester[T, R] {
	it.reader_func = reader
	return it
}

// Provide a solver function.
//
// N.B. all "ProvideSolver" calls ought to be made before providing expected outcomes
func (it *integration_tester[T, R]) ProvideSolver(solver func(T) (R, error)) *integration_tester[T, R] {
	if it.can_add_solvers {
		it.solvers = append(it.solvers, solver)
	} else {
		it.t.Error("IntegrationTester fatal error: Cannot add further solvers after providing test cases")
		it.t.FailNow()
	}
	return it
}

// Provide a test case i.e. an input filename and the expected outcomes to match results of each solver
//
// N.B. Number of expected outcomes ought to match the number of previously provided solvers.
// Also, no more solvers can be added after this method is invoked.
func (it *integration_tester[T, R]) AddTestCase(
	input_filename string,
	outcome expected_outcome[R],
	outcomes ...expected_outcome[R],
) *integration_tester[T, R] {
	it.can_add_solvers = false
	if len(outcomes)+1 != len(it.solvers) {
		it.t.Errorf(
			"IntegrationTester fatal error: number of solvers (%d) does not match the number of expected outcomes (%d)",
			len(it.solvers), len(outcomes),
		)
		it.t.FailNow()
	} else {
		test_case := integration_tester_test_case[R]{
			append([]expected_outcome[R]{outcome}, outcomes...),
			input_filename,
		}
		test_case.addPrefixes()
		it.test_cases = append(it.test_cases, test_case)
	}
	return it
}

// Returns true if everything is in order. Meaning:
//   - Equality function is provided
//   - Reader function is provided
//   - At least one test case is provided
func (it *integration_tester[T, R]) performChecksBeforeRunningTests() bool {
	if len(it.test_cases) == 0 {
		it.t.Error("There were no test cases provided; nothing to test")
		return false
	} else if it.reader_func == nil {
		it.t.Error("No reader function was provided; tests cannot be run")
		it.t.FailNow()
		return false
	} else if it.equality_func == nil {
		it.t.Error("No equality function was provided; tests cannot be run")
		it.t.FailNow()
		return false
	}
	return true
}

// Goes through all of the input filenames provided and on read input it runs each solver and verifies it matches the expected outcome
//
// N.B. equality function and reader function must be provided before calling this method
func (it *integration_tester[T, R]) RunIntegrationTests() *integration_tester[T, R] {
	if !it.performChecksBeforeRunningTests() {
		return it
	}
	for _, test_case := range it.test_cases {
		data, err := it.reader_func(test_case.filename)
		if err != nil {
			it.t.Errorf(
				`Integration Tester test case with input file "%s" failed: Reader produced the following error: "%s"`,
				test_case.filename, err.Error(),
			)
			continue
		}

		for i := 0; i < len(test_case.expected_outcomes); i++ {
			result, err := it.solvers[i](data)
			test_case.expected_outcomes[i].testOutcome(it.t, it.equality_func, result, err)
		}
	}
	return it
}
