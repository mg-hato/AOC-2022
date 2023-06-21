package models

import (
	c "aoc/common"
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

func CreateRucksacksEnvelope(rucksacks []Rucksack) c.Envelope[[]Rucksack] {
	return &RucksacksEnvelope{rucksacks}
}

func (be RucksacksEnvelope) String() string {
	return fmt.Sprintf("[%s]", strings.Join(be.rucksacks, ", "))
}

func RucksacksEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]Rucksack]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
