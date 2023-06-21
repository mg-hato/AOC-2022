package solver

import (
	c "aoc/common"
	m "aoc/d03/models"
	"fmt"
)

type badge_item_locator struct{}

// Constructor function for badge item locator
func BadgeItemLocator() ItemLocator {
	return &badge_item_locator{}
}

func (badge_item_locator) String() string {
	return "BadgeItemLocator"
}

func (badge_item_locator) locateItems(rucksacks []m.Rucksack) ([]rune, error) {
	// group rucksacks into groups of 3 and find any repeating items among each group (i.e. badge candidates)
	badge_candidates := c.Map(
		func(groupIndex int) []rune {
			return find_common_items(
				rucksacks[3*groupIndex],
				rucksacks[3*groupIndex+1],
				rucksacks[3*groupIndex+2],
			)
		},
		c.Range(0, len(rucksacks)/3),
	)

	// ensure that each group has exactly one badge, if not return with error
	if idx := c.IndexOf(badge_candidates, func(items []rune) bool { return len(items) != 1 }); idx != -1 {
		prefix := fmt.Sprintf(
			"Error while locating badges for rucksack group #%d (rucksacks %d-%d)",
			idx+1, 3*idx+1, 3*idx+3,
		)
		if len(badge_candidates[idx]) == 0 {
			return nil, fmt.Errorf("%s: no badge candidates (no item repeats among all 3 rucksacks)", prefix)
		} else {
			return nil, fmt.Errorf("%s: multiple badge candidates found: [%s]", prefix, string(badge_candidates[idx]))
		}
	}
	// All ok: for each group extract the badge
	return c.Map(func(badge []rune) rune { return badge[0] }, badge_candidates), nil
}
