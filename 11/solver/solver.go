package solver

import (
	c "aoc/common"
	m "aoc/d11/models"
	"sort"
)

func CalculateMonkeyBusiness(
	rounds_requested int,
	strategy WorryLevelAdjustmentStrategy,
) func(m.SolverInput) (int64, error) {
	return func(input m.SolverInput) (int64, error) {

		monkeys := c.Map(func(monkey m.Monkey) *m.Monkey { return &monkey }, input.Get())

		adjust := getAdjustmentFunction(
			strategy,
			c.Map(func(monkey *m.Monkey) int { return monkey.DivTest }, monkeys),
		)

		analysis_counter := make([]int, len(monkeys))

		for round_counter := 0; round_counter < rounds_requested; round_counter++ {
			for i, monkey := range monkeys {
				for _, item := range monkey.Items {
					adjusted_item := adjust(monkey.InspectionOP.Inspect(item))
					receiver := monkey.PerformDivisionTest(adjusted_item)
					monkeys[receiver].Receive(adjusted_item)
				}
				analysis_counter[i] += len(monkey.Items)
				monkey.Items = make([]int, 0)
			}
		}
		// Sort analysis counters from highest to lowest
		sort.Slice(analysis_counter, func(i, j int) bool { return analysis_counter[i] > analysis_counter[j] })
		return int64(analysis_counter[0]) * int64(analysis_counter[1]), nil
	}
}
