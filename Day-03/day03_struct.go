package main

import "fmt"

type ListOfContents struct {
	rucksacks []Rucksack
}

func (contents ListOfContents) String() string {
	return fmt.Sprintf("ListOfContents{ %s }", contents.rucksacks)
}

type Rucksack struct {
	items string
}

func (r Rucksack) FirstCompartment() string {
	return r.items[:len(r.items)/2]
}
func (r Rucksack) SecondCompartment() string {
	return r.items[len(r.items)/2:]
}

func (r Rucksack) String() string {
	return fmt.Sprintf(
		"R[(%s) (%s)]",
		r.FirstCompartment(),
		r.SecondCompartment(),
	)
}

// Get the priority of the passed item
func PriorityOf(item rune) int {
	switch {
	case 'a' <= item && item <= 'z':
		{
			return int(item - 'a' + 1)
		}
	default: // Must be A-Z
		{
			return int(item - 'A' + 27)
		}
	}
}
