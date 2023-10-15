package solver

import (
	c "aoc/common"
	m "aoc/d25/models"
	"fmt"
)

func SolveSnafu(envelope c.Envelope[[]string]) (string, error) {
	snafus := c.Map(m.SnafuNumberFromString, envelope.Get())
	sum := c.Sum(c.Map(m.SnafuNumber.ToInt, snafus))
	sum_snafu := m.SnafuNumberFromInt(sum)

	if sum != sum_snafu.ToInt() {
		fmt.Printf("snafu conversion seems off: expected %d vs actual %d", sum, sum_snafu.ToInt())
	}
	return sum_snafu.String(), nil
}
