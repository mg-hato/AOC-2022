package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func ReadListOfContents(filename string) (ListOfContents, error) {
	list := ListOfContents{[]Rucksack{}}
	file, e := os.Open(filename)

	// Try to open the file
	if e != nil {
		return list, e
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Empty-line regular expression
	empty_line_regexp := regexp.MustCompile("^ *$")

	// Valid-line regular expression
	valid_line := regexp.MustCompile("^([A-Za-z]{2})+$")

	var line_number int
	for scanner.Scan() {
		line_number++

		// Skip line if it is empty-line (contains only white-spaces)
		if empty_line_regexp.MatchString(scanner.Text()) {
			continue
		}

		if valid_line.MatchString(scanner.Text()) {
			list.rucksacks = append(list.rucksacks, Rucksack{scanner.Text()})
		} else {
			return list, error_BadInputLine(scanner.Text(), line_number)
		}
	}

	// If scanner left the loop because of an error
	if scanner.Err() != nil {
		return list, scanner.Err()
	}

	// If list of contents is not divisible by three, raise an error
	if len(list.rucksacks)%3 != 0 {
		return list, error_NumberOfRucksacksNotDivBy3(list)
	}

	return list, nil
}

// An error when a read line of input does not satisfy the regular expression
func error_BadInputLine(line string, line_number int) error {
	message := fmt.Sprintf("Bad line of input on line %d. Line: \"%s\"", line_number, line)
	return errors.New(message)
}

// Error when the total number of rucksacks is not divisible by 3
func error_NumberOfRucksacksNotDivBy3(list ListOfContents) error {
	message := fmt.Sprintf("Number of rucksacks is not divisible by 3 (number of rucksacks: %d)", len(list.rucksacks))
	return errors.New(message)
}
