package reader

import (
	c "aoc/common"
	m "aoc/d11/models"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type line_reader interface {
	// Process passed line and modify passed Monkey object.
	// Returns any potential error encountered
	ProcessLine(string, *m.Monkey) error
}

// 1st line: Monkey's identity number reader
type monkey_identity_line_reader struct {
	monkey_id_re *regexp.Regexp
}

func (reader monkey_identity_line_reader) ProcessLine(line string, monkey *m.Monkey) error {
	if submatches := reader.monkey_id_re.FindStringSubmatch(line); len(submatches) == 2 {
		monkey.MonkeyId, _ = strconv.Atoi(submatches[1])
		return nil
	}
	return errors.New("failed to extract monkey's ID")
}

func createMonkeyIdentityLineReader() line_reader {
	return monkey_identity_line_reader{
		monkey_id_re: regexp.MustCompile(`^ *Monkey (\d+): *$`),
	}
}

// 2nd line: Monkey's starting items
type monkey_starting_items_line_reader struct {
	monkey_starting_items_re *regexp.Regexp
}

func (reader monkey_starting_items_line_reader) ProcessLine(line string, monkey *m.Monkey) error {
	if submatches := reader.monkey_starting_items_re.FindStringSubmatch(line); len(submatches) == 2 {
		monkey.Items = c.Map(
			func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			},
			c.Filter(
				func(s string) bool { return len(s) > 0 },
				c.Map(strings.TrimSpace, strings.Split(submatches[1], ",")),
			),
		)
		return nil
	}
	return errors.New("failed to extract monkey's starting items")
}

func createMonkeyStartingItemsLineReader() line_reader {
	return monkey_starting_items_line_reader{
		monkey_starting_items_re: regexp.MustCompile(`^ *Starting items: *(\d+ *(?:, *\d+ *)*)? *$`),
	}
}

// 3rd line: Monkey's operation
type monkey_operation_line_reader struct {
	monkey_operation_re *regexp.Regexp
}

func (reader monkey_operation_line_reader) ProcessLine(line string, monkey *m.Monkey) error {
	if submatches := reader.monkey_operation_re.FindStringSubmatch(line); len(submatches) == 4 {
		monkey.InspectionOP = m.IOP(
			makeOperand(submatches[1]),
			makeOperator(submatches[2]),
			makeOperand(submatches[3]),
		)
		return nil
	}
	return errors.New("failed to extract monkey's operation")
}

func makeOperand(s string) m.Operand {
	if s == "old" {
		return m.Old()
	} else {
		i, _ := strconv.Atoi(s)
		return m.Num(i)
	}
}

func makeOperator(s string) m.Operator {
	switch s {
	case "*":
		return m.Multiplication{}
	default:
		return m.Addition{}
	}
}

func createMonkeyOperationLineReader() line_reader {
	return monkey_operation_line_reader{
		monkey_operation_re: regexp.MustCompile(`^ *Operation: new *= *(old|\d+) *(\*|\+) *(old|\d+) *$`),
	}
}

// 4th line: Monkey's test
type monkey_test_line_reader struct {
	monkey_test_re *regexp.Regexp
}

func (reader monkey_test_line_reader) ProcessLine(line string, monkey *m.Monkey) error {
	if submatches := reader.monkey_test_re.FindStringSubmatch(line); len(submatches) == 2 {
		i, _ := strconv.Atoi(submatches[1])
		monkey.DivTest = i
		return nil
	}
	return errors.New("failed to extract monkey's test")
}

func createMonkeyTestLineReader() line_reader {
	return monkey_test_line_reader{
		monkey_test_re: regexp.MustCompile(`^ *Test: divisible by (\d+) *$`),
	}
}

// 5th & 6th line: Monkey's if-clause
type monkey_if_clause_line_reader struct {
	monkey_if_clause_re *regexp.Regexp
	clause_boolean      bool
}

func (reader monkey_if_clause_line_reader) ProcessLine(line string, monkey *m.Monkey) error {
	if submatches := reader.monkey_if_clause_re.FindStringSubmatch(line); len(submatches) == 2 {
		value, _ := strconv.Atoi(submatches[1])
		if reader.clause_boolean {
			monkey.OnTrue = value
		} else {
			monkey.OnFalse = value
		}
		return nil
	}
	return fmt.Errorf("failed to extract monkey's if-%v throwing action", reader.clause_boolean)
}

func createMonkeyIfClauseLineReader(clause bool) line_reader {
	return monkey_if_clause_line_reader{
		monkey_if_clause_re: regexp.MustCompile(fmt.Sprintf("^ *If %v: throw to monkey (\\d+) *$", clause)),
		clause_boolean:      clause,
	}
}
