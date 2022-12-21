package functional

import "testing"

func TestArrayEqual_Integers(t *testing.T) {

	inputs := []struct {
		lhs, rhs        []int
		expectedOutcome bool
	}{
		{lhs: []int{10, 5, 100, 51}, rhs: []int{10, 5, 100, 51}, expectedOutcome: true},
		{lhs: []int{10, 5}, rhs: []int{10, 5, 100, 51}, expectedOutcome: false},
		{lhs: []int{10, 5, 100, 51}, rhs: []int{10, 5}, expectedOutcome: false},
		{lhs: []int{10, 5, 100, 51, 0}, rhs: []int{10, 5, 100, 51}, expectedOutcome: false},
		{lhs: []int{}, rhs: []int{}, expectedOutcome: true},
		{lhs: []int{-7}, rhs: []int{-7}, expectedOutcome: true},
		{lhs: []int{}, rhs: []int{0}, expectedOutcome: false},
	}

	for test_number, input := range inputs {
		if actual := ArrayEqual(input.lhs, input.rhs); actual != input.expectedOutcome {
			t.Errorf("Test input #%d failed. ArrayEqual(%v, %v) returned %v, but %v was expected\n",
				test_number+1,
				input.lhs, input.rhs,
				actual, input.expectedOutcome)
		}
	}
}

func TestArrayEqual_Strings(t *testing.T) {

	inputs := []struct {
		lhs, rhs        []string
		expectedOutcome bool
	}{
		{lhs: []string{"One", "Two", "String"}, rhs: []string{"One", "Two", "String"}, expectedOutcome: true},
		{lhs: []string{"One", "Two", "AlmOst"}, rhs: []string{"One", "Two", "Almost"}, expectedOutcome: false}, // 3rd diff
		{lhs: []string{"One", "Two"}, rhs: []string{"OneTwo"}, expectedOutcome: false},
		{lhs: []string{"", ""}, rhs: []string{"", ""}, expectedOutcome: true},
		{lhs: []string{}, rhs: []string{}, expectedOutcome: true},
		{lhs: []string{}, rhs: []string{""}, expectedOutcome: false},
	}

	for test_number, input := range inputs {
		if actual := ArrayEqual(input.lhs, input.rhs); actual != input.expectedOutcome {
			t.Errorf("Test input #%d failed. ArrayEqual(%v, %v) returned %v, but %v was expected\n",
				test_number+1,
				input.lhs, input.rhs,
				actual, input.expectedOutcome)
		}
	}
}
