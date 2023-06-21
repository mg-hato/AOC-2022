package reader

import (
	m "aoc/d02/models"
	"aoc/reading"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type strategy_guide_reader struct {
	err    error
	rounds []m.Round

	line_number int

	empty_re *regexp.Regexp
	round_re *regexp.Regexp

	left_symbol_mapping  map[string]m.LeftSymbol
	right_symbol_mapping map[string]m.RightSymbol
}

// Constructor function for Strategy Guide Reader
func StategyGuideReader() reading.ReaderAoC2022[m.SolverInput] {
	return &strategy_guide_reader{
		err: nil,

		rounds: make([]m.Round, 0),

		line_number: 0,

		empty_re: regexp.MustCompile("^ *$"),
		round_re: regexp.MustCompile("^ *([ABCabc]) +([XYZxyz]) *$"),

		left_symbol_mapping: map[string]m.LeftSymbol{
			"A": m.A,
			"B": m.B,
			"C": m.C,
		},

		right_symbol_mapping: map[string]m.RightSymbol{
			"X": m.X,
			"Y": m.Y,
			"Z": m.Z,
		},
	}
}

func (sgr strategy_guide_reader) Error() error {
	return sgr.err
}

func (sgr strategy_guide_reader) PerformFinalValidation() error {
	if len(sgr.rounds) == 0 {
		return errors.New("Error: Strategy Guide is empty")
	}
	return nil
}

func (sgr strategy_guide_reader) Done() bool {
	return sgr.err != nil
}

func (sgr *strategy_guide_reader) ProvideLine(line string) {
	sgr.line_number++
	if sgr.empty_re.MatchString(line) {
		return
	}

	if matches := sgr.round_re.FindStringSubmatch(line); len(matches) == 3 {
		sgr.rounds = append(sgr.rounds, m.Round{
			Left:  sgr.left_symbol_mapping[strings.ToUpper(matches[1])],
			Right: sgr.right_symbol_mapping[strings.ToUpper(matches[2])],
		})
	} else {
		sgr.err = fmt.Errorf(`Error: could not interpret line #%d: "%s"`, sgr.line_number, line)
	}
}

func (sgr strategy_guide_reader) FinishAndGetInputData() m.SolverInput {
	return m.CreateRoundsEnvelope(sgr.rounds)
}
