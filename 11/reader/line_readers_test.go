package reader

import (
	m "aoc/day11/models"
	f "aoc/functional"
	"aoc/testers"
	"testing"
)

type line_reader_test_input[T any] struct {
	line     string
	expected T
}

func lrti[T any](line string, expected T) line_reader_test_input[T] {
	return line_reader_test_input[T]{line: line, expected: expected}
}

func TestD11_MonkeyIdentityLineReader(t *testing.T) {
	line_reader := createMonkeyIdentityLineReader()
	monkey := new(m.Monkey)

	testers.TestThat([]line_reader_test_input[int]{
		lrti("Monkey 12:", 12),
		lrti("  Monkey 0:  ", 0),
		lrti("    Monkey 2021:", 2021),
	}, func(input line_reader_test_input[int]) {
		testers.AssertNoError(t, line_reader.ProcessLine(input.line, monkey))
		testers.AssertEqual(t, monkey.MonkeyId, input.expected)
	})
}

func TestD11_MonkeyStartingItemsLineReader(t *testing.T) {
	line_reader := createMonkeyStartingItemsLineReader()
	monkey := new(m.Monkey)

	testers.TestThat([]line_reader_test_input[[]int]{
		lrti("Starting items: ", []int{}),
		lrti("  Starting items: 1,  2  , 5", []int{1, 2, 5}),
		lrti("     Starting items:   1001 ,10  ,  990  ,1    ", []int{1001, 10, 990, 1}),
	}, func(input line_reader_test_input[[]int]) {
		testers.AssertNoError(t, line_reader.ProcessLine(input.line, monkey))
		testers.AssertEqualWithEqFunc(t, monkey.Items, input.expected, f.ArrayEqual[int])
	})
}

func TestD11_MonkeyOpearationLineReader(t *testing.T) {
	line_reader := createMonkeyOperationLineReader()
	monkey := new(m.Monkey)
	eq_func := func(lhs, rhs m.InspectionOperation) bool { return lhs == rhs }

	testers.TestThat([]line_reader_test_input[m.InspectionOperation]{
		lrti(
			"Operation: new = old + old",
			m.IOP(m.Old(), m.Add(), m.Old()),
		),
		lrti(
			"Operation: new=1001*99",
			m.IOP(m.Num(1001), m.Mult(), m.Num(99)),
		),
		lrti(
			"  Operation: new=  303   *old   ",
			m.IOP(m.Num(303), m.Mult(), m.Old()),
		),
		lrti(
			"  Operation: new=  old+  1",
			m.IOP(m.Old(), m.Add(), m.Num(1)),
		),
	}, func(input line_reader_test_input[m.InspectionOperation]) {
		testers.AssertNoError(t, line_reader.ProcessLine(input.line, monkey))
		testers.AssertEqualWithEqFunc(t, monkey.InspectionOP, input.expected, eq_func)
	})
}

func TestD11_MonkeyTestLineReader(t *testing.T) {
	line_reader := createMonkeyTestLineReader()
	monkey := new(m.Monkey)

	testers.TestThat([]line_reader_test_input[int]{
		lrti("    Test: divisible by 12  ", 12),
		lrti("Test: divisible by 1", 1),
		lrti("Test: divisible by 1000    ", 1000),
	}, func(input line_reader_test_input[int]) {
		testers.AssertNoError(t, line_reader.ProcessLine(input.line, monkey))
		testers.AssertEqual(t, monkey.DivTest, input.expected)
	})
}
func TestD11_MonkeyIfClauseLineReader_true(t *testing.T) {
	line_reader := createMonkeyIfClauseLineReader(true)
	monkey := new(m.Monkey)

	testers.TestThat([]line_reader_test_input[int]{
		lrti("If true: throw to monkey 10", 10),
		lrti("If true: throw to monkey 7  ", 7),
		lrti("   If true: throw to monkey 11 ", 11),
	}, func(input line_reader_test_input[int]) {
		testers.AssertNoError(t, line_reader.ProcessLine(input.line, monkey))
		testers.AssertEqual(t, monkey.OnTrue, input.expected)
	})
}

func TestD11_MonkeyIfClauseLineReader_false(t *testing.T) {
	line_reader := createMonkeyIfClauseLineReader(false)
	monkey := new(m.Monkey)

	testers.TestThat([]line_reader_test_input[int]{
		lrti("If false: throw to monkey 10", 10),
		lrti("If false: throw to monkey 7  ", 7),
		lrti("   If false: throw to monkey 11 ", 11),
	}, func(input line_reader_test_input[int]) {
		testers.AssertNoError(t, line_reader.ProcessLine(input.line, monkey))
		testers.AssertEqual(t, monkey.OnFalse, input.expected)
	})
}
