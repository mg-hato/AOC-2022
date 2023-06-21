package reader

import (
	c "aoc/common"
	m "aoc/d04/models"
	"aoc/reading"
	"fmt"
	"regexp"
	"strconv"
)

type section_assignments_list_reader struct {
	err error

	line_number int

	assignment_re *regexp.Regexp

	assignments []m.SectionAssignmentPair
}

// Constructor function for section assignments list reader
func SectionAssignmentsListReader() reading.ReaderAoC2022[m.SolverInput] {
	return &section_assignments_list_reader{
		err: nil,

		line_number: 0,

		assignment_re: regexp.MustCompile(`^ *(\d+)-(\d+) *, *(\d+)-(\d+) *$`),

		assignments: make([]m.SectionAssignmentPair, 0),
	}
}

func (salr section_assignments_list_reader) Error() error {
	return salr.err
}

func (salr section_assignments_list_reader) PerformFinalValidation() error {
	return nil
}

func (salr section_assignments_list_reader) Done() bool {
	return salr.Error() != nil
}

func (salr *section_assignments_list_reader) ProvideLine(line string) {
	salr.line_number++

	matches := salr.assignment_re.FindStringSubmatch(line)
	if len(matches) != 5 {
		salr.err = fmt.Errorf(`Error: cannot interpret line #%d: "%s"`, salr.line_number, line)
		return
	}

	// Make assignment pair
	values := c.Map(func(s string) int { i, _ := strconv.Atoi(s); return i }, matches[1:])
	assignment := m.MakeSAP(values[0], values[1], values[2], values[3])

	// Ensure assignment pair is valid i.e. each elf's coverage is the correct way around
	if assignment.First.Left > assignment.First.Right || assignment.Second.Left > assignment.Second.Right {
		salr.err = fmt.Errorf(
			`Error: line #%d assignment pair has one or more coverages in invalid order: %s`,
			salr.line_number,
			assignment,
		)
		return
	}
	// All ok: add assignment
	salr.assignments = append(salr.assignments, assignment)
}

func (salr section_assignments_list_reader) FinishAndGetInputData() m.SolverInput {
	return m.CreateSectionAssignmentPairsEnvelope(salr.assignments)
}
