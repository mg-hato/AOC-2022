package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadListOfAssignmentPairs(filename string) ([]AssignmentPair, error) {

	// Try to open the file
	file, e := os.Open(filename)
	if e != nil {
		return nil, e
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var line_number int
	assignmentPairs := []AssignmentPair{}

	for scanner.Scan() {
		line_number++

		// Parse line
		line := scanner.Text()
		if ap, err := parseAssignmentPair(line); err != nil {
			file.Close()
			message := fmt.Sprintf("Error encountered parsing a line #%d. Line: \"%s\"\nError: %s", line_number, line, err.Error())
			return nil, errors.New(message)
		} else {
			assignmentPairs = append(assignmentPairs, ap)
		}

	}
	file.Close()

	// Check if any error
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return assignmentPairs, nil
}

// Parse assignment pair and returns any error detected
func parseAssignmentPair(line string) (AssignmentPair, error) {
	var ap AssignmentPair

	sections := strings.Split(line, ",")
	if len(sections) != 2 {
		return ap, error_NotTwoSectionsPerLine(sections)
	}

	// Parse sections of the first elf
	if s, err := parseSections(sections[0]); err != nil {
		return ap, err
	} else {
		ap.first = s
	}

	// Parse sections of the second elf
	if s, err := parseSections(sections[1]); err != nil {
		return ap, err
	} else {
		ap.second = s
	}

	return ap, nil
}

// Parses sections and returns any error detected
func parseSections(sectionDesc string) (SectionRange, error) {
	var sections SectionRange
	numbers := strings.Split(sectionDesc, "-")

	// Check that exactly two numbers are given
	if len(numbers) != 2 {
		return sections, error_NotTwoNumbersPerSectionDescription(numbers)
	}

	// Check that on the left a number is given
	if n, err := strconv.Atoi(numbers[0]); err == nil {
		sections.left = n
	} else {
		return sections, error_SectionRangeNotANumber(numbers[0], "left")
	}

	// Check that on the right a number is given
	if n, err := strconv.Atoi(numbers[1]); err == nil {
		sections.right = n
	} else {
		return sections, error_SectionRangeNotANumber(numbers[1], "right")
	}

	// Check that left-right numbers define a proper range
	if sections.left > sections.right {
		return sections, error_BadRange(sections)
	}

	return sections, nil
}

// A line is expected to describe a pair of sections, where section descriptions are connected with a ','
//
// Day 04 specification: "the Elves pair up and make a big list of the section assignments for each pair"
func error_NotTwoSectionsPerLine(sections []string) error {
	message := fmt.Sprintf("A line of input is expected to have exactly 2 sections described connected with a comma (','), but it has %d. They are: %v", len(sections), sections)
	return errors.New(message)
}

// A section description is expected to be described with two numbers connected with a '-'
//
// Day 04 specification: "... was assigned sections 2-4 (sections 2, 3, and 4) ..."
func error_NotTwoNumbersPerSectionDescription(numbers []string) error {
	message := fmt.Sprintf("A section description is expected to have exactly 2 numbers connected with a dash ('-'), but it has %d. They are: %v", len(numbers), numbers)
	return errors.New(message)
}

// A section description is expected to be described with two numbers connected with a '-'
//
// Day 04 specification: "... was assigned sections 2-4 (sections 2, 3, and 4) ..."
func error_SectionRangeNotANumber(input string, leftOrRight string) error {
	message := fmt.Sprintf(
		"A section description is expected to have exactly 2 numbers connected with a dash ('-'). On the %s of '-' the input provided was not a number: %s",
		leftOrRight, input,
	)
	return errors.New(message)
}

// A section should define a valid range (left-right with left <= right)
//
// Day 04 specification: "... was assigned sections 2-4 (sections 2, 3, and 4) ..."
func error_BadRange(s SectionRange) error {
	message := fmt.Sprintf("A section read is not defining a valid range. Read: %d-%d", s.left, s.right)
	return errors.New(message)
}
