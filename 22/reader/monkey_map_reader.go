package reader

import (
	c "aoc/common"
	m "aoc/d22/models"
	"aoc/reading"
	"regexp"
	"strconv"
)

type monkey_map_reader struct {
	err         error
	line_number int
	done_flag   bool

	fields       []m.Field
	instructions []m.Instruction

	map_re *regexp.Regexp

	instructions_re *regexp.Regexp
}

func MonkeyMapReader() reading.ReaderAoC2022[c.Envelope[c.Pair[[]m.Field, []m.Instruction]]] {
	return &monkey_map_reader{
		fields:       make([]m.Field, 0),
		instructions: make([]m.Instruction, 0),

		map_re: regexp.MustCompile(`^[ \.#]*$`),

		instructions_re: regexp.MustCompile(`^ *(\d+|[LR])+ *$`),
	}
}

func (mmr *monkey_map_reader) Error() error {
	return mmr.err
}

func (mmr monkey_map_reader) PerformFinalValidation() error {
	for _, validation_func := range []func(monkey_map_reader) error{
		validate_at_least_one_dot,
		validate_continuity,
	} {
		validation_error := validation_func(mmr)
		if validation_error != nil {
			return validation_error
		}
	}
	return nil
}

func (mmr monkey_map_reader) Done() bool {
	return mmr.err != nil || mmr.done_flag
}

func (mmr *monkey_map_reader) ProvideLine(line string) {
	mmr.line_number++
	switch {
	case mmr.map_re.MatchString(line):
		mmr.handle_map_line(line)
	case mmr.instructions_re.MatchString(line):
		mmr.handle_instructions_line(line)
	default:
		mmr.err = could_not_interpret_line_reading_error(mmr.line_number, line)
	}
}

func (mmr *monkey_map_reader) FinishAndGetInputData() c.Envelope[c.Pair[[]m.Field, []m.Instruction]] {
	return m.CreateFieldsAndInstructionsEnvelope(mmr.fields, mmr.instructions)
}

// Internal receiver functions

func (mmr *monkey_map_reader) handle_map_line(line string) {
	fields := make([]m.Field, 0)
	for column_num, r := range line {
		ft, err := m.TryParseFieldType(r)
		if err != nil {
			mmr.err = err
			return
		}
		// Add field
		fields = append(fields, m.Field{
			FType:  ft,
			Row:    mmr.line_number,
			Column: column_num + 1,
		})
	}

	fields = c.Filter(func(field m.Field) bool { return field.FType != m.Blank }, fields)
	mmr.fields = append(mmr.fields, fields...)
}

func (mmr *monkey_map_reader) handle_instructions_line(line string) {
	for _, instr := range regexp.MustCompile(`\d+|[LR]`).FindAllString(line, -1) {
		switch instr {
		case "R":
			mmr.instructions = append(mmr.instructions, m.CreateTurnInstruction(m.Right))
		case "L":
			mmr.instructions = append(mmr.instructions, m.CreateTurnInstruction(m.Left))
		default:
			i, _ := strconv.Atoi(instr)
			mmr.instructions = append(mmr.instructions, m.CreateMoveInstruction(i))
		}
	}
	mmr.done_flag = true
}
