package solver

import (
	c "aoc/common"
	m "aoc/d01/models"
	"fmt"
	"sort"
)

func CalorySumOfTop(n int) func(m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {
		var sums []int = c.Map(c.Sum[int], input.Get())
		if len(sums) < n {
			return 0, fmt.Errorf(
				`error: In solver function "CalorySumOfTop(n = %d)" the calory list passed is of length %d`,
				n, len(sums),
			)
		}
		sort.Slice(sums, func(i, j int) bool { return sums[i] > sums[j] })
		return c.Sum(c.Take(n, sums)), nil
	}
}
