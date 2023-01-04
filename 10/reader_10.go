package main

import (
	"aoc/reading"
	"fmt"
	"regexp"
	"strconv"
)

type InstructionsReader struct {
	err  error
	data []Instruction

	line_number int
	noop_regexp *regexp.Regexp
	addx_regexp *regexp.Regexp
}

// CONSTRUCTOR

func NewInstructionReader() reading.ReaderAoC2022[[]Instruction] {
	return &InstructionsReader{
		noop_regexp: regexp.MustCompile("^noop$"),
		addx_regexp: regexp.MustCompile(`^addx (-?\d+)$`),
	}
}

// PUBLIC METHODS

func (ir *InstructionsReader) Error() error {
	return ir.err
}

func (ir *InstructionsReader) PerformFinalValidation() error {
	return nil
}

func (ir *InstructionsReader) Done() bool {
	return ir.Error() != nil
}

func (ir *InstructionsReader) ProvideLine(line string) {
	ir.line_number++
	if ir.noop_regexp.MatchString(line) {
		ir.data = append(ir.data, NewNoop())
	} else if addx_parts := ir.addx_regexp.FindStringSubmatch(line); len(addx_parts) == 2 {
		arg, _ := strconv.Atoi(addx_parts[1])
		ir.data = append(ir.data, NewAddx(arg))
	} else {
		ir.err = fmt.Errorf("Error: Unknown instruction on line %d: \"%s\"", ir.line_number, line)
	}

}

func (ir *InstructionsReader) FinishAndGetInputData() []Instruction {
	return ir.data
}
