package solver

import (
	m "aoc/day03/models"
	f "aoc/functional"
)

// Finds common items across all provided rucksacks
func find_common_items(rucksack m.Rucksack, rucksacks ...m.Rucksack) []rune {
	convert_to_set := func(r string) map[rune]bool {
		return f.AssociateWith([]rune(r), func(rune) bool { return true })
	}
	candidate_items := convert_to_set(rucksack)
	for i := 0; i < len(rucksacks); i++ {
		current_rucksack := convert_to_set(rucksacks[i])
		for _, item := range f.GetKeys(candidate_items) {
			if !current_rucksack[item] {
				delete(candidate_items, item)
			}
		}
	}

	return f.GetKeys(candidate_items)
}
