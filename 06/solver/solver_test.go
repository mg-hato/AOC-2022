package solver

import (
	ts "aoc/testers"
	"testing"
)

func TestD06_SolverTest(t *testing.T) {
	ts.SolverTesterForComparableResults[string, int](t).
		ProvideSolver(FindPositionOfTheFirstMarker(4)).
		ProvideSolver(FindPositionOfTheFirstMarker(14)).
		AddTestCase("abc", ts.ExpectError[int]("no marker found"), ts.ExpectError[int]("no marker found")).
		AddTestCase("abcd", ts.ExpectResult(4), ts.ExpectError[int]("no marker found")).
		AddTestCase("abcdefghijklmn", ts.ExpectResult(4), ts.ExpectResult(14)).
		RunSolverTests()
}

func TestD06_FrequencyCounterTest(t *testing.T) {
	fc := make_frequency_counter()
	ts.AssertEqual(t, fc.different_count, 0)

	fc.addElement('a')
	fc.addElement('a')
	fc.addElement('a')
	fc.addElement('a')
	ts.AssertEqual(t, fc.different_count, 1) // 4 'a'

	fc.removeElement('a')
	fc.removeElement('a')
	fc.removeElement('a')
	fc.addElement('x')
	ts.AssertEqual(t, fc.different_count, 2) // 1 'x' & 1 'a'

	fc.removeElement('a')
	ts.AssertEqual(t, fc.different_count, 1) // only 1 'x'
}
