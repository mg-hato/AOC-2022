package reader

import (
	"aoc/day11/models"
	"aoc/envelope"
	f "aoc/functional"
	"aoc/reading"
	"errors"
	"fmt"
	"regexp"
)

type monkey_graph_reader struct {
	line_readers  []line_reader
	empty_line_re *regexp.Regexp

	reader_index, line_number int

	monkeys []models.Monkey
	monkey  *models.Monkey

	e error
}

// Constructor of MonkeyGraphReader
func MonkeyGraphReader() reading.ReaderAoC2022[envelope.Envelope[[]models.Monkey]] {
	return &monkey_graph_reader{
		line_readers: []line_reader{
			createMonkeyIdentityLineReader(),
			createMonkeyStartingItemsLineReader(),
			createMonkeyOperationLineReader(),
			createMonkeyTestLineReader(),
			createMonkeyIfClauseLineReader(true),
			createMonkeyIfClauseLineReader(false),
		},
		reader_index:  0,
		line_number:   0,
		empty_line_re: regexp.MustCompile(`^ *$`),

		monkeys: make([]models.Monkey, 0),
		monkey:  new(models.Monkey),

		e: nil,
	}
}

// ReaderAoC2022 methods implementations

func (mgr *monkey_graph_reader) Error() error {
	return mgr.e
}

func (mgr *monkey_graph_reader) PerformFinalValidation() error {

	// Ensure that there are at least two monkeys
	if len(mgr.monkeys) < 2 {
		return fmt.Errorf(
			"Error: Monkey Graph needs at least 2 monkeys to calculate monkey business value, but this graph has %d monkey(s)",
			len(mgr.monkeys),
		)
	}

	// Ensure that the monkey IDs are from 0 incrementing
	for expectedId, monkey := range mgr.monkeys {
		if expectedId != monkey.MonkeyId {
			return errors.New(
				"Error: monkey IDs are not in incrementing order from 0 onwards",
			)
		}
	}

	// Ensure that no monkey performs divisibility test with 0
	for id, monkey := range mgr.monkeys {
		if monkey.DivTest == 0 {
			return fmt.Errorf("Error: Monkey %d performs divisibility test with 0", id)
		}
	}

	// Ensure that no monkey passes items to itself (self-loop)
	for id, monkey := range mgr.monkeys {
		if monkey.OnFalse == id || monkey.OnTrue == id {
			return fmt.Errorf("Error: Monkey %d should not pass items to itself (self-loop)", id)
		}
	}

	// Ensure that each monkey passes items to another valid monkey (i.e. monkey in scope)
	for id, monkey := range mgr.monkeys {
		if !f.InRange(monkey.OnFalse, 0, len(mgr.monkeys)) || !f.InRange(monkey.OnTrue, 0, len(mgr.monkeys)) {
			return fmt.Errorf("Error: Monkey %d can pass items to a monkey outside of the expected range [0, %d)", id, len(mgr.monkeys))
		}
	}
	return nil
}

func (mgr *monkey_graph_reader) Done() bool {
	return mgr.Error() != nil
}

func (mgr *monkey_graph_reader) ProvideLine(line string) {
	mgr.line_number++

	if !mgr.empty_line_re.MatchString(line) {
		if e := mgr.line_readers[mgr.reader_index].ProcessLine(line, mgr.monkey); e != nil {
			mgr.e = fmt.Errorf("Error on line #%d: %s. Line: %s", mgr.line_number, e.Error(), line)
		}

		// Next line reader index
		mgr.reader_index = (mgr.reader_index + 1) % len(mgr.line_readers)

		// If all readers have finished save the current monkey & refresh it
		if mgr.reader_index == 0 {
			mgr.monkeys = append(mgr.monkeys, *mgr.monkey)
			mgr.monkey = new(models.Monkey)
		}
	}
}

func (mgr *monkey_graph_reader) FinishAndGetInputData() envelope.Envelope[[]models.Monkey] {
	return models.CreateMonkeysEnvelopeWith(mgr.monkeys)
}
