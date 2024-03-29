package models

import (
	c "aoc/common"
	"aoc/testers"
	"testing"
)

func TestD03_RucksacksEnvelope(t *testing.T) {
	envelope := CreateRucksacksEnvelope([]Rucksack{"abcABc", "xAzXxZ", "PQRSSpqA"})
	rucksacks := envelope.Get()
	rucksacks[0] = "tBtw"

	testers.AssertEqualWithEqFunc(t, envelope.Get(), []Rucksack{"abcABc", "xAzXxZ", "PQRSSpqA"}, c.ArrayEqual[string])
}
