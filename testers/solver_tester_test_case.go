package testers

import "fmt"

type solver_tester_test_case[T, R any] struct {
	expected_outcomes []expected_outcome[R]
	input_data        T
}

func (sttc *solver_tester_test_case[T, R]) addPrefixes() *solver_tester_test_case[T, R] {
	for i := 0; i < len(sttc.expected_outcomes); i++ {
		sttc.expected_outcomes[i].setPrefix(fmt.Sprintf(
			`Solver tester test case #%d for input data %v has failed`,
			i+1, sttc.input_data,
		))
	}
	return sttc
}
