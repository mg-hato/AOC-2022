package reader

import (
	m "aoc/d10/models"
	"aoc/reading"
	"fmt"
	"regexp"
	"strconv"
)

type instructions_reader struct {
	err  error
	data []m.Instruction

	line_number int
	noop_re     *regexp.Regexp
	addx_re     *regexp.Regexp
	empty_re    *regexp.Regexp
}

func InstructionsReader() reading.ReaderAoC2022[m.SolverInput] {
	return &instructions_reader{
		noop_re:  regexp.MustCompile("^ *noop *$"),
		addx_re:  regexp.MustCompile(`^ *addx +(0|-?[1-9]\d*) *$`),
		empty_re: regexp.MustCompile(`^ *$`),
	}
}

func (ir *instructions_reader) Error() error {
	return ir.err
}

func (ir *instructions_reader) PerformFinalValidation() error {
	return nil
}

func (ir *instructions_reader) Done() bool {
	return ir.Error() != nil
}

func (ir *instructions_reader) ProvideLine(line string) {
	ir.line_number++
	addx_submatches := ir.addx_re.FindStringSubmatch(line)

	switch {
	case addx_submatches != nil:
		v, _ := strconv.Atoi(addx_submatches[1])
		ir.data = append(ir.data, m.Addx(v))
	case ir.noop_re.MatchString(line):
		ir.data = append(ir.data, m.Noop())
	case ir.empty_re.MatchString(line):
	default:
		ir.err = fmt.Errorf(`error on line #%d: could not interpret "%s"`, ir.line_number, line)
	}
}

func (ir *instructions_reader) FinishAndGetInputData() m.SolverInput {
	return m.InstructionsEnvelope(ir.data)
}
