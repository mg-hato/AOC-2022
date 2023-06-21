package solver

import (
	c "aoc/common"
	m "aoc/d03/models"
)

func SumItemPriorities(locator ItemLocator) func(m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {
		items, err := locator.locateItems(input.Get())
		return c.Sum(c.Map(get_item_priority, items)), err
	}
}
