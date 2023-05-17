package solver

import (
	m "aoc/day03/models"
	f "aoc/functional"
	"fmt"
)

type compartment_duplicate_item_locator struct{}

// Constructor function for compartment duplicate item locator
func CompartmentDuplicateItemLocator() ItemLocator {
	return &compartment_duplicate_item_locator{}
}

func (compartment_duplicate_item_locator) String() string {
	return "CompartmentDuplicateItemLocator"
}

func (compartment_duplicate_item_locator) locateItems(rucksacks []m.Rucksack) ([]rune, error) {
	// split each rucksack into its two compartments and find any common items
	repeating_item_candidates := f.Map(func(r m.Rucksack) []rune {
		mid := len(r) / 2
		return find_common_items(r[:mid], r[mid:])
	}, rucksacks)

	// we expect exactly one item to be duplicate between both compartments, if that is not the case return with error
	if idx := f.IndexOf(repeating_item_candidates, func(items []rune) bool { return len(items) != 1 }); idx != -1 {
		prefix := fmt.Sprintf("Error while locating repeating items in compartments of rucksack #%d", idx+1)
		if len(repeating_item_candidates[idx]) == 0 {
			return nil, fmt.Errorf("%s: no repeating items", prefix)
		} else {
			return nil, fmt.Errorf(
				"%s: multiple repeating items found: [%s]",
				prefix, string(repeating_item_candidates[idx]),
			)
		}
	}
	// All ok: for each rucksack extract the repeating item
	return f.Map(func(repeating_item []rune) rune { return repeating_item[0] }, repeating_item_candidates), nil
}
