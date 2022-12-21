package functional

import (
	"fmt"
	"strings"
	"testing"
)

func TestMapEqual_IntStr(t *testing.T) {
	tests := []struct {
		lhs, rhs        map[int]string
		expectedOutcome bool
	}{
		{
			lhs:             map[int]string{1: "1", 10: "$%", 7: "Seven"},
			rhs:             map[int]string{7: "Seven", 1: "1", 10: "$%"},
			expectedOutcome: true,
		},
		{
			lhs:             map[int]string{1: "1", 10: "$%", 7: "Seven"},
			rhs:             map[int]string{7: "seven", 1: "1", 10: "$%"}, // value that 7 maps to is different
			expectedOutcome: false,
		},
		{
			lhs:             map[int]string{1: "1", 10: "$%", 7: "Seven"},
			rhs:             map[int]string{7: "Seven", 1: "1", 11: "$%"}, // different keys
			expectedOutcome: false,
		},
		{
			lhs:             map[int]string{},
			rhs:             map[int]string{},
			expectedOutcome: true,
		},
		{
			lhs:             map[int]string{1: "1"},
			rhs:             map[int]string{1: "1"},
			expectedOutcome: true,
		},
		{
			lhs:             map[int]string{1: "1"},
			rhs:             map[int]string{1: "1", 2: "2"},
			expectedOutcome: false,
		},
		{
			lhs:             map[int]string{1: "1"},
			rhs:             map[int]string{1: "1", 2: "2"},
			expectedOutcome: false,
		},
	}

	for test_number, test := range tests {
		if result := MapEqual(test.lhs, test.rhs); result != test.expectedOutcome {
			msg := strings.Join([]string{
				fmt.Sprintf("Test number #%d failed. Tested MapEqual(%v, %v)\n", test_number+1, test.lhs, test.rhs),
				fmt.Sprintf("\tActual outcome: %v\n", result),
				fmt.Sprintf("\tExpected outcome: %v\n", test.expectedOutcome),
			}, "")
			t.Error(msg)
		}
	}
}
