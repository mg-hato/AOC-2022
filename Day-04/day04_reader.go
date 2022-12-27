package main

import (
	"aoc/functional"
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadListOfAssignmentPairs(filename string) ([]AssignmentPair, error) {

	// Try to open the file
	file, e := os.Open(filename)
	if e != nil {
		return nil, e
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Regular expression for a valid line
	valid_line_regexp := regexp.MustCompile("^(\\d+)-(\\d+),(\\d+)-(\\d+)$")

	var line_number int
	assignmentPairs := []AssignmentPair{}

	for scanner.Scan() {
		line_number++

		// If the current line does not match the expect input format: return error
		if !valid_line_regexp.MatchString(scanner.Text()) {
			return nil, error_BadInputLine(line_number, scanner.Text())
		}

		ap := getAssignmentPair(valid_line_regexp.FindStringSubmatch(scanner.Text())[1:])

		if !ap.validateAssignmentPair() {
			return nil, error_BadSectionRange(line_number, ap)
		}

		assignmentPairs = append(assignmentPairs, ap)
	}

	// Check if any error
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return assignmentPairs, nil
}

// Get assignment-pair from list of a verified list of strings
func getAssignmentPair(nums []string) AssignmentPair {
	sections := functional.Map(
		func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}, nums,
	)

	return AssignmentPair{
		first:  SectionRange{left: sections[0], right: sections[1]},
		second: SectionRange{left: sections[2], right: sections[3]},
	}
}

func (pair AssignmentPair) validateAssignmentPair() bool {
	return pair.first.validateSectionRange() && pair.second.validateSectionRange()
}

func (sr SectionRange) validateSectionRange() bool {
	return sr.left <= sr.right
}

// A section should define a valid range (left-right with left <= right)
func error_BadSectionRange(line_number int, ap AssignmentPair) error {
	message := fmt.Sprintf("A section-range read is not defining a valid range (line %d). Assignment pair: %s", line_number, ap)
	return errors.New(message)
}

func error_BadInputLine(line_number int, line string) error {
	message := fmt.Sprintf("Bad input line on line %d. Line is \"%s\"", line_number, line)
	return errors.New(message)
}
