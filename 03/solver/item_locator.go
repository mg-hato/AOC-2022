package solver

import m "aoc/day03/models"

type ItemLocator interface {
	String() string
	locateItems([]m.Rucksack) ([]rune, error)
}
