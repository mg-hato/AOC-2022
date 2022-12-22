package main

import (
	. "aoc/functional"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadListOfContents(filename string) (*ListOfContents, error) {
	file, e := os.Open(filename)

	// Try to open the file
	if e != nil {
		file.Close()
		return nil, e
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	contents := ListOfContents{[]Rucksack{}}

	var lineNumber int = 0
	for scanner.Scan() {
		lineNumber++

		// If error reading the line, return
		if e := scanner.Err(); e != nil {
			file.Close()
			return nil, e
		}

		rucksack, problems := validateAndReturnRucksack(scanner.Text())

		// If any error is detected during the validation of the read rucksack, raise them and stop reading
		if len(problems) != 0 {

			// Close the file
			file.Close()

			// Turn errors into enumerated list of error messages
			listOfErrorMessages := Map(
				func(p Pair[int, error]) string {
					return fmt.Sprintf("\t%d. %s", p.First, p.Second.Error())
				},
				EnumerateWithFirstIndex(problems, 1),
			)

			// Combine them and present them nicely
			combinedErrorMessage := fmt.Sprintf(
				"%d error(s) encountered when validating input of line %d:\n%s",
				len(listOfErrorMessages),
				lineNumber,
				strings.Join(listOfErrorMessages, "\n"),
			)

			return nil, errors.New(combinedErrorMessage)
		}

		contents.rucksacks = append(contents.rucksacks, rucksack)
	}

	file.Close()
	return &contents, error_NumberOfRucksacksNotDivBy3(&contents)
}

// Create rucksack and return an array of errors (if any)
func validateAndReturnRucksack(items string) (Rucksack, []error) {
	rucksack := Rucksack{items}
	problems := []error{}

	if len(items)%2 != 0 {
		problems = append(problems, error_OddNumberOfItems(len(items)))
	}

	for _, item := range items {
		if !(('a' <= item && item <= 'z') || ('A' <= item && item <= 'Z')) {
			problems = append(problems, error_InvalidItem(item))
		}
	}

	return rucksack, problems
}

// Error when a rucksack has an odd number of items
//
// (Day 03 part 1 specification: "A given rucksack always has the same number of items in each of its two compartments")
func error_OddNumberOfItems(items_count int) error {
	message := fmt.Sprintf("Rucksack is expected to have an even number of items, but the number is odd: %d", items_count)
	return errors.New(message)
}

// Error when an item is not valid i.e. it is not a character a-z or A-Z
//
// (Day 03 part 1 specification: "every item can be converted to a priority: ... Lowercase item types `a` through `z` ... Uppercase item types `A` through `Z` ...")
func error_InvalidItem(item rune) error {
	message := fmt.Sprintf("Item \"%c\" does not satisfy the specification (a-z and A-Z)", item)
	return errors.New(message)
}

// Error when the total number of rucksacks is not divisible by 3
//
// (Day 03 part 2 specification: "For Safety Elves (Rucksacks) are divided into groups of three")
func error_NumberOfRucksacksNotDivBy3(contents *ListOfContents) error {
	if size := len(contents.rucksacks); size%3 != 0 {
		message := fmt.Sprintf("Number of rucksacks is not divisible by 3 (number of rucksacks: %d)", size)
		return errors.New(message)
	}
	return nil
}
