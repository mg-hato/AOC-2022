package main

import (
	. "aoc/functional"
	"aoc/reading"
	"errors"
	"fmt"
	"regexp"
)

type ForestReader struct {
	err    error    // any error
	forest [][]byte // Heights of trees

	line_number int

	// Regular expression for a line of input (1-digit tree-heights)
	tree_regexp *regexp.Regexp
}

// CONSTRUCTOR

func NewForestReader() reading.ReaderAoC2022[[][]byte] {
	return &ForestReader{
		err:         nil,
		forest:      [][]byte{},
		tree_regexp: regexp.MustCompile("^(\\d+)$"),
	}
}

// PUBLIC METHODS

func (fr *ForestReader) Error() error {
	return fr.err
}

func (fr *ForestReader) PerformFinalValidation() error {

	// Ensure that the input data forms a grid
	// i.e. that all input lines have the same numbers of digits
	lengths := Map(func(line []byte) int { return len(line) }, fr.forest)
	if !All(func(l int) bool { return lengths[0] == l }, lengths) {
		return errors.New(fmt.Sprintf(
			"Error: The input digits do not form correct (rectangular) grid "+
				"as they are of different sizes. Sizes are: %v",
			lengths,
		))
	}

	return nil
}

func (fr *ForestReader) Done() bool {
	return fr.err != nil
}

func (fr *ForestReader) ProvideLine(line string) {
	fr.line_number++

	if subs := fr.tree_regexp.FindStringSubmatch(line); len(subs) == 2 {
		fr.forest = append(fr.forest, Map(func(b byte) byte { return b - '0' }, []byte(subs[1])))
	} else {
		fr.err = errors.New(fmt.Sprintf("Error: Bad input on line %d (expected numbers only). Line: \"%s\"", fr.line_number, line))
	}
}

func (fr *ForestReader) FinishAndGetInputData() [][]byte {
	return fr.forest
}
