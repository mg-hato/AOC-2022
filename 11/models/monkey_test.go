package models

import (
	"aoc/functional"
	"aoc/testers"
	"testing"
)

func test_monkey() Monkey {
	return Monkey{
		DivTest: 33,
		OnTrue:  1,
		OnFalse: 2,
		Items:   []int{1, 2, 3},
	}
}

func TestD11_MonkeyReceiveTest(t *testing.T) {
	monkey := test_monkey()
	monkey.Receive(15)
	monkey.Receive(100)
	testers.AssertEqualWithEqFunc(t, monkey.Items, []int{1, 2, 3, 15, 100}, functional.ArrayEqual[int])
}

func TestD11_MonkeyPerformDivisionTestTest(t *testing.T) {
	monkey := test_monkey()
	testers.AssertEqual(t, monkey.PerformDivisionTest(10), 2)
	testers.AssertEqual(t, monkey.PerformDivisionTest(99), 1)
	testers.AssertEqual(t, monkey.PerformDivisionTest(100), 2)
	testers.AssertEqual(t, monkey.PerformDivisionTest(33), 1)
	testers.AssertEqual(t, monkey.PerformDivisionTest(0), 1)
}
