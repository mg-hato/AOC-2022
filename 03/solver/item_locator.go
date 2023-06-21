package solver

import m "aoc/d03/models"

type ItemLocator interface {
	String() string
	locateItems([]m.Rucksack) ([]rune, error)
}
