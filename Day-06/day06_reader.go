package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func ReadDatastream(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Input is only
	scanner.Scan()

	// If any error during reading
	if scanner.Err() != nil {
		return "", scanner.Err()
	}

	if regexp.MustCompile("^[a-z]+$").MatchString(scanner.Text()) {
		return scanner.Text(), nil
	} else {
		return "", error_BadInput()
	}
}

func error_BadInput() error {
	message := fmt.Sprint("Error: Input is only expected to be single string consisting of only lower-case characters.")
	return errors.New(message)
}
