package solver

import (
	c "aoc/common"
	m "aoc/d20/models"
	ts "aoc/testers"
	"testing"
)

func simple_decrypt(arr []int) int64 {

	// special case: 0
	if len(arr) == 1 {
		return 0
	}

	// First: initial index; Second: value
	e_arr := c.Enumerate(arr)
	for i := 0; i < len(arr); i++ {
		position := c.IndexOf(e_arr, func(pair c.Pair[int, int]) bool { return pair.First == i })
		selected := e_arr[position]
		e_arr = append(e_arr[:position], e_arr[position+1:]...)
		position = (position + selected.Second) % len(e_arr)
		if position < 0 {
			position += len(e_arr)
		}
		left, right := e_arr[:position], e_arr[position:]
		e_arr = c.FlatMap(c.Identity[[]c.Pair[int, int]], [][]c.Pair[int, int]{left, {selected}, right})
	}

	zero_position := c.IndexOf(e_arr, func(pair c.Pair[int, int]) bool { return pair.Second == 0 })

	return c.Sum(c.Map(func(offset int) int64 {
		return int64(e_arr[(zero_position+offset)%len(e_arr)].Second)
	}, []int{1000, 2000, 3000}))
}

func TestD20_Decrypt(t *testing.T) {
	solver_tester := ts.SolverTesterForComparableResults[m.SolverInput, int64](t).ProvideSolver(Decrypt(1, int64(1)))
	for _, input := range [][]int{
		{0},
		{0, 1, 1},
		{1, 2, 3, 4, 0},
		c.Range(0, 113),
		c.Range(-7, 11),
		{-6, 0, 1, 2, 15, -9, -3, 1},
	} {
		envelope := m.EncryptedFileEnvelope(input)
		expected_result := simple_decrypt(envelope.Get())
		solver_tester.AddTestCase(envelope, ts.ExpectResult(expected_result))
	}
	solver_tester.RunSolverTests()
}
