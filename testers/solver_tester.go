package testers

import (
	"testing"
)

type solver_tester[T, R any] struct {
	t *testing.T

	solvers []func(T) (R, error)

	can_add_solvers bool

	test_cases []solver_tester_test_case[T, R]

	eq_func func(R, R) bool
}

// Solver Tester constructor function
func SolverTester[T, R any](t *testing.T) *solver_tester[T, R] {
	return &solver_tester[T, R]{
		t: t,

		solvers:         make([]func(T) (R, error), 0),
		test_cases:      make([]solver_tester_test_case[T, R], 0),
		can_add_solvers: true,
	}
}

// Solver Tester constructor function for comparable results
func SolverTesterForComparableResults[T any, R comparable](t *testing.T) *solver_tester[T, R] {
	return SolverTester[T, R](t).
		ProvideEqualityFunction(func(lhs, rhs R) bool { return lhs == rhs })
}

// Provide an equality function for results produced by the solver(s)
func (st *solver_tester[T, R]) ProvideEqualityFunction(eq_func func(R, R) bool) *solver_tester[T, R] {
	st.eq_func = eq_func
	return st
}

// Provide a solver function.
//
// N.B. all "ProvideSolver" calls ought to be made before providing expected outcomes
func (st *solver_tester[T, R]) ProvideSolver(solver func(T) (R, error)) *solver_tester[T, R] {
	if st.can_add_solvers {
		st.solvers = append(st.solvers, solver)
	} else {
		st.t.Error("SolverTester fatal error: Cannot add further solvers after providing test cases")
		st.t.FailNow()
	}
	return st
}

// Provide a test case i.e. an input data with expected outcomes for each solver
//
// N.B. Number of expected outcomes ought to match the number of previously provided solvers.
// Also, no more solver can be added after this method is invoked.
func (st *solver_tester[T, R]) AddTestCase(
	data T,
	outcome expected_outcome[R],
	outcomes ...expected_outcome[R],
) *solver_tester[T, R] {
	st.can_add_solvers = false
	if len(outcomes)+1 != len(st.solvers) {
		st.t.Errorf(
			"SolverTester fatal error: number of solvers (%d) does not match the number of expected outcomes (%d)",
			len(st.solvers), len(outcomes),
		)
		st.t.FailNow()
	} else {
		test_case := solver_tester_test_case[T, R]{
			append([]expected_outcome[R]{outcome}, outcomes...),
			data,
		}
		test_case.addPrefixes()
		st.test_cases = append(st.test_cases, test_case)
	}
	return st
}

func (st *solver_tester[T, R]) performChecksBeforeRunningTests() bool {
	if len(st.test_cases) == 0 {
		st.t.Error("There were no test cases provided; nothing to test")
		return false
	} else if st.eq_func == nil {
		st.t.Error("No equality function was provided; tests cannot be run")
		st.t.FailNow()
		return false
	}
	return true
}

func (st *solver_tester[T, R]) RunSolverTests() *solver_tester[T, R] {
	if !st.performChecksBeforeRunningTests() {
		return st
	}
	for _, test_case := range st.test_cases {
		for i := 0; i < len(st.solvers); i++ {
			result, err := st.solvers[i](test_case.input_data)
			test_case.expected_outcomes[i].testOutcome(st.t, st.eq_func, result, err)
		}
	}
	return st
}
