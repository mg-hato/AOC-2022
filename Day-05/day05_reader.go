package main

import (
	. "aoc/functional"
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadRearrangementPlan(filename string) (RearrangementPlan, error) {
	var plan RearrangementPlan

	file, err := os.Open(filename)
	if err != nil {
		return plan, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var line_number int
	plan.stacks, err = readStacksDrawing(scanner, &line_number)
	if err != nil {
		return plan, err
	}

	plan.moves, err = readRearrangementInstructions(scanner, &line_number)
	if err != nil {
		return plan, err
	}

	// Validate that all move-instructions are valid (i.e. the source stack_id and the destination stack_id are within range of [1..len(stacks)])
	if err := validateMoveInstructions(plan.moves, len(plan.stacks)); err != nil {
		return plan, err
	}

	return plan, nil
}

// Reads the first part of the input representing the starting box-stacks
func readStacksDrawing(scanner *bufio.Scanner, line_number *int) ([]string, error) {
	lines := []string{}

	// Firstly: scan lines until a line describing stack ids is found a.k.a. "stack-id line"
	// Also, validate each line of boxes (on any level) for correctness of their format
	// and return error if validation fails at any level

	// Regular expression for 1 box or for empty box (3-whitespaces)
	box_regexp := "( {3}|\\[[A-Z]\\])"

	// Regular expression for boxes (1 or more) separated by single-whitespace (with possibly multiple trailing white-spaces)
	boxes_regexp := fmt.Sprintf("((%s( %s)*) *)", box_regexp, box_regexp)

	// Regular expression that a line describing any level of boxes should satisfy:
	// 1. it is either fully empty (0 or more white-spaces)
	// 2. OR describes some aligned boxes and has 0 or more trailing white-spaces
	box_line_regexp := regexp.MustCompile(fmt.Sprintf("^(%s| *)$", boxes_regexp))

	var stack_ids []int
	for scanner.Scan() {
		*line_number++
		line := scanner.Text()
		stack_ids = tryParseStackIds(line)

		if stack_ids == nil {
			// This line is not stack-id line, so it must be a "box-line"
			// Validate that the format of the boxes is correct
			if !box_line_regexp.MatchString(line) {
				return nil, error_boxLineValidationFailed(*line_number, line)
			}
			lines = append(lines, line)
		} else {
			// This line is stack-id line; break out of the loop
			break
		}
	}

	// If scanner left the Scan because of the error
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	// If stack-id line not found
	if stack_ids == nil {
		return nil, error_stackIdLineNotFound()
	}

	// If validation of stack-ids fails
	if validation_err := validateStackIds(stack_ids); validation_err != nil {
		return nil, validation_err
	}

	lines = Reverse(lines)
	stacks := []string{}
	errs := []error{}

	// Collect all boxes
	for i := 1; i <= len(stack_ids); i++ {
		boxes := Map(func(s string) rune { return getBox(s, i) }, lines)

		// Validate that no "floating-boxes" are present
		if e := validateThatNoBoxIsFloating(i, boxes); e != nil {
			errs = append(errs, e)
		}

		// Remove the "empty-boxes" from the top (if any)
		boxes = Filter(func(box rune) bool { return box != ' ' }, boxes)
		stacks = append(stacks, string(boxes))
	}

	// If any errors are encountered in the previous step, format and return them
	if len(errs) != 0 {
		return nil, error_formatMultipleErrors(errs)
	}

	return stacks, nil
}

// Reads the second part of the input: The rearrangement instructions
func readRearrangementInstructions(scanner *bufio.Scanner, line_number *int) ([]Move, error) {
	instructions := []Move{}

	empty_line_regexp := regexp.MustCompile("^ *$")
	move_regexp := regexp.MustCompile("^move (\\d+) from (\\d+) to (\\d+)$")

	// errors to be collected
	errs := []error{}

	for scanner.Scan() {
		*line_number++

		// Allow empty lines: ignore them
		// If not an empty line, it must match the move-instruction syntax (otherwise, return an error)
		if empty_line_regexp.MatchString(scanner.Text()) {
			continue
		} else if submatches := move_regexp.FindStringSubmatch(scanner.Text()); submatches != nil && len(submatches) == 4 {
			instructions = append(instructions, interpretAsMove(submatches))
		} else {
			errs = append(errs, error_moveInstructionParsingError(scanner.Text(), *line_number))
		}
	}

	// If any move-instructions failed to parse, format those errors and return them
	if len(errs) != 0 {
		return nil, error_formatMultipleErrors(errs)
	}

	// If scanner encountered any errors
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return instructions, nil
}

func interpretAsMove(submatches []string) Move {
	numbers := Map(func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}, submatches[1:]) // Skip the first (whole line as a match)
	return Move{qty: numbers[0], source: numbers[1], destination: numbers[2]}
}

// Parses the and returns stack ids from it iff the line matches regexp for the stack-id line
func tryParseStackIds(line string) []int {

	// Check whether the line matches the regular expression for the stack-id line
	// Note: stack-id line contains only non-negative integers separated by (perhaps repeated) white-spaces
	can_parse, _ := regexp.MatchString("^ *([0-9]+ *)+$", line)

	if can_parse {
		// Extract a splice of integers from the line (ignoring all the repeated white spaces)
		return Map(
			func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			},
			Filter(
				func(s string) bool { return len(s) != 0 },
				strings.Split(line, " "),
			))
	}

	return nil
}

// Returns an error if validation of stack ids fails
func validateStackIds(stack_ids []int) error {
	// Stack ids must be (in-order) numbers: 1, 2, ... len(stack_ids)

	// Create array of expected stack ids
	expected_stack_ids := make([]int, len(stack_ids))
	for i := 0; i < len(stack_ids); i++ {
		expected_stack_ids[i] = i + 1
	}

	if !ArrayEqual(stack_ids, expected_stack_ids) {
		return error_stackIdsValidationFailed(stack_ids, expected_stack_ids)
	}

	return nil
}

// Retrieves the box (or lack of it i.e. white-space)
func getBox(line string, stack_id int) rune {

	// Position of a box w.r.t. stack index examples:
	// 1 -> 1; 2 -> 5; 3 -> 9; ...
	box_position := (stack_id-1)*4 + 1

	if box_position >= len(line) {
		return ' '
	} else {
		return rune(line[box_position])
	}
}

// Validate that there are no "floating boxes"
//
// i.e. return an error if there is a "physical" (non-white-space) box under which is empty-space
//
// Note: A box on the ground (bottom) cannot be a "floating box"
func validateThatNoBoxIsFloating(stack_id int, boxes []rune) error {
	var below_is_empty bool
	for _, box := range boxes {
		if below_is_empty && box != ' ' {
			return error_floatingBox(stack_id)
		}
		below_is_empty = box == ' '
	}
	return nil
}

// Validate move instructions for source and destination stack_ids
func validateMoveInstructions(moves []Move, stack_id_size int) error {

	// move-instruction is valid if it moves from and to a valid stack-id
	// a stack-id is valid if it exists, meaning it is a value between 1 and stack_id_size (inclusive)
	stack_id_ok := func(i int) bool {
		return 1 <= i && i <= stack_id_size
	}

	// Validate that for every move-instruction both source & destination are valid
	errs := FlatMap(func(pair Pair[int, Move]) []error {
		errs := []error{}
		if !stack_id_ok(pair.Second.source) {
			errs = append(errs, error_moveInstructionInvalidSourceOrDestination("source", pair.First, pair.Second.source, stack_id_size))
		}
		if !stack_id_ok(pair.Second.destination) {
			errs = append(errs, error_moveInstructionInvalidSourceOrDestination("destination", pair.First, pair.Second.destination, stack_id_size))
		}
		return errs
	}, EnumerateWithFirstIndex(moves, 1))

	// return formatted error if any errors are encountered
	if len(errs) != 0 {
		return error_formatMultipleErrors(errs)
	}

	return nil
}

func error_formatMultipleErrors(errs []error) error {
	listed_errors := Map(
		func(p Pair[int, error]) string {
			return fmt.Sprintf("\n\t%d. %s", p.First, p.Second.Error())
		},
		EnumerateWithFirstIndex(errs, 1),
	)
	return errors.New(fmt.Sprintf("%d error(s) encountered:%s", len(listed_errors), strings.Join(listed_errors, "")))
}

func error_boxLineValidationFailed(line_number int, line string) error {
	message := fmt.Sprintf(
		"Error: Box line validation failed on line %d. Line validated: \"%s\"",
		line_number, line,
	)
	return errors.New(message)
}

func error_moveInstructionParsingError(line string, line_number int) error {
	message := fmt.Sprintf("Error: Could not parse a move instruction on line %d. Line: \"%s\"", line_number, line)
	return errors.New(message)
}

func error_floatingBox(stack_id int) error {
	message := fmt.Sprintf(
		"Error: Found a floating box(es) (it is \"flying mid-air\") in stack %d",
		stack_id,
	)
	return errors.New(message)
}

func error_stackIdLineNotFound() error {
	message := "Error: Could not find a line that describes stack-ids"
	return errors.New(message)
}

func error_stackIdsValidationFailed(actual_stack_ids, expected_stack_ids []int) error {
	message := fmt.Sprintf(
		"Error: The line identified to be the stack-id line has ids: %v, but expected ids are: %v",
		actual_stack_ids, expected_stack_ids,
	)
	return errors.New(message)
}

func error_moveInstructionInvalidSourceOrDestination(field_name string, move_id, field_value, stack_id_size int) error {
	message := fmt.Sprintf(
		"Error: Move instruction #%d has invalid field \"%s\" equal to %d (expected values are 1 to %d inclusive)",
		move_id, field_name, field_value, stack_id_size,
	)
	return errors.New(message)
}
