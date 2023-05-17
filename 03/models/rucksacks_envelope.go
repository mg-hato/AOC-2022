package models

import (
	e "aoc/envelope"
	"aoc/functional"
	"fmt"
	"strings"
)

type RucksacksEnvelope struct {
	rucksacks []Rucksack
}

func (be RucksacksEnvelope) Get() []Rucksack {
	new_rucksacks := make([]Rucksack, len(be.rucksacks))
	copy(new_rucksacks, be.rucksacks)
	return new_rucksacks
}

func CreateRucksacksEnvelope(rucksacks []Rucksack) e.Envelope[[]Rucksack] {
	return &RucksacksEnvelope{rucksacks}
}

func (be RucksacksEnvelope) String() string {
	return fmt.Sprintf("[%s]", strings.Join(be.rucksacks, ", "))
}

func RucksacksEnvelopeEqualityFunction(lhs, rhs e.Envelope[[]Rucksack]) bool {
	return functional.ArrayEqual(lhs.Get(), rhs.Get())
}
