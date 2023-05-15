package solver

import (
	m "aoc/day01/models"
	e "aoc/envelope"
	f "aoc/functional"
	"fmt"
	"sort"
)

func CalorySumOfTop(n int) func(e.Envelope[m.CaloryList]) (int, error) {
	return func(envelope e.Envelope[m.CaloryList]) (int, error) {
		var sums []int = f.Map(f.Sum[int], envelope.Get())
		if len(sums) < n {
			return 0, fmt.Errorf(
				`error: In solver function "CalorySumOfTop(n = %d)" the calory list passed is of length %d`,
				n, len(sums),
			)
		}
		sort.Slice(sums, func(i, j int) bool { return sums[i] > sums[j] })
		return f.Sum(f.Take(n, sums)), nil
	}
}
