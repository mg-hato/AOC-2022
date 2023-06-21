package solver

import (
	c "aoc/common"
	m "aoc/d03/models"
)

// Finds common items across all provided rucksacks
func find_common_items(rucksack m.Rucksack, rucksacks ...m.Rucksack) []rune {
	convert_to_set := func(r string) map[rune]bool {
		return c.AssociateWith([]rune(r), func(rune) bool { return true })
	}
	candidate_items := convert_to_set(rucksack)
	for i := 0; i < len(rucksacks); i++ {
		current_rucksack := convert_to_set(rucksacks[i])
		for _, item := range c.GetKeys(candidate_items) {
			if !current_rucksack[item] {
				delete(candidate_items, item)
			}
		}
	}

	return c.GetKeys(candidate_items)
}
