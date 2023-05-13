package testers

import "fmt"

type reader_tester_test_case[T any] struct {
	expected_outcome expected_outcome[T]
	filename         string
}

func (rttc *reader_tester_test_case[T]) addPrefix() *reader_tester_test_case[T] {
	rttc.expected_outcome.setPrefix(fmt.Sprintf(
		`Reader tester test case for input file "%s" has failed`,
		rttc.filename,
	))
	return rttc
}
