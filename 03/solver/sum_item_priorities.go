package solver

import (
	m "aoc/day03/models"
	e "aoc/envelope"
	f "aoc/functional"
)

func SumItemPriorities(locator ItemLocator) func(e.Envelope[[]m.Rucksack]) (int, error) {
	return func(envelope e.Envelope[[]m.Rucksack]) (int, error) {
		items, err := locator.locateItems(envelope.Get())
		return f.Sum(f.Map(get_item_priority, items)), err
	}
}
