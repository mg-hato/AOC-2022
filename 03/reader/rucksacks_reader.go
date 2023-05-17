package reader

import (
	m "aoc/day03/models"
	e "aoc/envelope"
	"aoc/reading"
	"fmt"
	"regexp"
)

type rucksacks_reader struct {
	err error

	line_number int

	rucksack_re *regexp.Regexp

	rucksacks []m.Rucksack
}

// Constructor function for rucksacks reader
func RucksacksReader() reading.ReaderAoC2022[e.Envelope[[]m.Rucksack]] {
	return &rucksacks_reader{
		err: nil,

		line_number: 0,

		rucksack_re: regexp.MustCompile(`^([A-Za-z]+)$`),

		rucksacks: make([]m.Rucksack, 0),
	}
}

func (rr rucksacks_reader) Error() error {
	return rr.err
}

func (rr rucksacks_reader) PerformFinalValidation() error {
	if len(rr.rucksacks)%3 != 0 {
		return fmt.Errorf(
			"Error: the number of elves/rucksacks (%d) cannot be split into groups of three",
			len(rr.rucksacks),
		)
	}
	return nil
}

func (rr rucksacks_reader) Done() bool {
	return rr.Error() != nil
}

func (rr *rucksacks_reader) ProvideLine(line string) {
	rr.line_number++

	// Check it matches RE
	matches := rr.rucksack_re.FindStringSubmatch(line)
	if len(matches) != 2 {
		rr.err = fmt.Errorf(`Error: cannot interpret line #%d: "%s"`, rr.line_number, line)
		return
	}
	// Ensure rucksack contains even number of items
	rucksack := matches[1]
	if len(rucksack)%2 != 0 {
		rr.err = fmt.Errorf(
			`Error: on line #%d, the rucksack must contain an even number of items. Number of items: %d; Rucksack: "%s"`,
			rr.line_number,
			len(rucksack),
			rucksack,
		)
		return
	}
	// All ok: add rucksack
	rr.rucksacks = append(rr.rucksacks, matches[1])
}

func (rr rucksacks_reader) FinishAndGetInputData() e.Envelope[[]m.Rucksack] {
	return m.CreateRucksacksEnvelope(rr.rucksacks)
}
