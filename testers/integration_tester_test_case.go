package testers

import (
	"fmt"
)

type integration_tester_test_case[R any] struct {
	expected_outcomes []expected_outcome[R]
	filename          string
}

func (ittc *integration_tester_test_case[R]) addPrefixes() *integration_tester_test_case[R] {
	for i := 0; i < len(ittc.expected_outcomes); i++ {
		ittc.expected_outcomes[i].setPrefix(fmt.Sprintf(
			`Integration tester test case #%d for input file "%s" has failed`,
			i+1, ittc.filename,
		))
	}
	return ittc
}
