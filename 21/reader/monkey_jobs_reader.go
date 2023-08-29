package reader

import (
	c "aoc/common"
	m "aoc/d21/models"
	"aoc/reading"
	"regexp"
	"strconv"
)

type monkey_jobs_reader struct {
	line_number int
	err         error

	monkey_jobs []c.Pair[string, m.MonkeyJob]
	monkey_ids  map[string]int

	empty_re         *regexp.Regexp
	simple_job_re    *regexp.Regexp
	operation_job_re *regexp.Regexp

	number_re *regexp.Regexp
}

func MonkeyJobsReader() reading.ReaderAoC2022[m.SolverInput] {
	return &monkey_jobs_reader{
		monkey_jobs: make([]c.Pair[string, m.MonkeyJob], 0),
		monkey_ids:  make(map[string]int),

		empty_re:         regexp.MustCompile(`^ *$`),
		simple_job_re:    regexp.MustCompile(`^ *([A-Za-z]+) *: *([A-Za-z]+|-?[0-9]+) *$`),
		operation_job_re: regexp.MustCompile(`^ *([A-Za-z]+) *: *([A-Za-z]+|-?[0-9]+) *([\+\-\*\/]) *([A-Za-z]+|-?[0-9]+) *$`),

		number_re: regexp.MustCompile(`^-?[0-9]+$`),
	}
}

func (mjr monkey_jobs_reader) Error() error {
	return mjr.err
}

func (mjr monkey_jobs_reader) PerformFinalValidation() error {
	for _, validation_func := range []func(monkey_jobs_reader) error{
		validate_existence_of_given_id("root"),
		validate_existence_of_given_id("humn"),
		validate_no_self_reference,
		validate_no_unknown_dependents,
	} {
		if validation_err := validation_func(mjr); validation_err != nil {
			return validation_err
		}
	}
	return nil
}

func (mjr monkey_jobs_reader) Done() bool {
	return mjr.Error() != nil
}

func (mjr *monkey_jobs_reader) ProvideLine(line string) {
	mjr.line_number++
	match_empty := mjr.empty_re.MatchString(line)
	match_simple_job := mjr.simple_job_re.FindStringSubmatch(line)
	match_operation_job := mjr.operation_job_re.FindStringSubmatch(line)
	switch {
	case match_empty:
	case match_simple_job != nil:
		mjr.handle_simple_job(match_simple_job[1], match_simple_job[2])
	case match_operation_job != nil:
		mjr.handle_operation_job(
			match_operation_job[1],
			match_operation_job[2],
			match_operation_job[3],
			match_operation_job[4],
		)
	default:
		mjr.err = bad_line_reading_error(mjr.line_number, line)
	}
}

func (mjr monkey_jobs_reader) FinishAndGetInputData() m.SolverInput {
	return m.CreateMonkeyJobsEnvelope(mjr.monkey_jobs...)
}

// Other private methods

// Handles job that does not include operation
func (mjr *monkey_jobs_reader) handle_simple_job(id, operand string) {
	if !mjr.check_id_uniqueness(id) {
		return
	}
	interpreted_operand := mjr.interpret_operand(operand)
	if _, is_number := interpreted_operand.(m.Number); !is_number {
		mjr.err = simple_job_is_not_a_number(mjr.line_number)
		return
	}
	job := m.CreateSingleJob(mjr.interpret_operand(operand))
	mjr.monkey_jobs = append(mjr.monkey_jobs, c.MakePair(id, job))
}

// Handles a job that includes arithmetic operation=
func (mjr *monkey_jobs_reader) handle_operation_job(id, lhs, op, rhs string) {
	if !mjr.check_id_uniqueness(id) {
		return
	}
	operation, op_err := m.TryParseOperation(op)
	if op_err != nil {
		mjr.err = unknown_operation_reading_error(mjr.line_number, op_err)
		return
	}
	inerpreted_lhs := mjr.interpret_operand(lhs)
	interpreted_rhs := mjr.interpret_operand(rhs)
	are_identifiers := c.All(func(o m.Operand) bool {
		_, is_identifier := o.(m.Identifier)
		return is_identifier
	}, []m.Operand{inerpreted_lhs, interpreted_rhs})

	if !are_identifiers {
		mjr.err = operation_job_must_have_only_identifiers_reading_error(mjr.line_number)
		return
	}
	job := m.CreateTwoOperandJob(inerpreted_lhs, operation, interpreted_rhs)
	mjr.monkey_jobs = append(mjr.monkey_jobs, c.MakePair(id, job))
}

// Verifies whether ID has been already used and iff not returns true
func (mjr *monkey_jobs_reader) check_id_uniqueness(id string) bool {
	if id_line_number, not_unique := mjr.monkey_ids[id]; not_unique {
		mjr.err = repeated_id_reading_error(mjr.line_number, id, id_line_number)
		return false
	}
	mjr.monkey_ids[id] = mjr.line_number
	return true
}

// Interprets operand; whether Number or Identifier
func (mjr *monkey_jobs_reader) interpret_operand(operand string) m.Operand {
	if mjr.number_re.MatchString(operand) {
		n, _ := strconv.ParseInt(operand, 10, 64)
		return m.CreateNumber(n)
	} else {
		return m.CreateIdentifier(operand)
	}
}
