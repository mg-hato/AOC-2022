package main

import (
	. "aoc/functional"
)

// Finds the sum of character-priorities that are fetched by `runesFetcher`
func SumOfPriorities(lc ListOfContents, runesFetcher func(ListOfContents) []rune) int {
	return Sum(Map(PriorityOf, runesFetcher(lc)))
}

// Find runes that are repeated in both compartments (part 1 of day 03)
func FindRepeatedItems(lc ListOfContents) []rune {
	return Map(
		func(r Rucksack) rune {
			firstCompartment := AssociateWith(
				[]rune(r.FirstCompartment()),
				func(_ rune) bool { return true },
			)
			return Filter(
				func(c rune) bool { return firstCompartment[c] },
				[]rune(r.SecondCompartment()),
			)[0]
		},
		lc.rucksacks,
	)
}

// Find badges of groups
func FindGroupBadges(lc ListOfContents) []rune {

	// Turn array of rucksacks into array of 3-rucksack groups
	groups := GetValues(GroupBy(
		Enumerate(lc.rucksacks),
		func(p Pair[int, Rucksack]) int { return p.First / 3 },
		func(p Pair[int, Rucksack]) Rucksack { return p.Second },
	))

	return Map(
		func(group []Rucksack) rune {
			var badge rune

			// Go through all items of each rucksack in the group
			// Using bitwise-operations create a bitmap s.t. i-th bit represents membership of i-th rucksack
			// (e.g. if item 'k' is a member of 3rd rucksack, 3rd bit of bitmap['k'] will be 1)
			bitmap := map[rune]int{}
			var bit int = 1
			for _, rucksack := range group {
				ForEach(
					func(c rune) { bitmap[c] = bitmap[c] | bit },
					[]rune(rucksack.items),
				)
				bit <<= 1 // Once done with current rucksack, left-shift the bit (equivalent to multiplication with 2)
			}

			// Once done, an item belonging to all three rucksacks will have first 3 bits set to 1 i.e. the number will be equal to 7 (1+2+4)
			for c, bit := range bitmap {
				if bit == 7 {
					badge = c
				}
			}
			return badge
		},
		groups,
	)
}
