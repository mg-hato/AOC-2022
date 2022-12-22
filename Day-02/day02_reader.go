package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadStrategy(filename string) (*EncryptedStrategyGuide, error) {
	file, e := os.Open(filename)

	// Try to open the file
	if e != nil {
		file.Close()
		return nil, e
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	strategy := EncryptedStrategyGuide{rounds: []Round{}}

	var lineNumber int = 0
	for scanner.Scan() {
		lineNumber++

		// If error reading the line, return
		if e := scanner.Err(); e != nil {
			file.Close()
			return nil, e
		}

		round, e := parseRound(strings.Split(scanner.Text(), " "))

		// If error during parsing the line, return
		if e != nil {
			file.Close()
			msg := fmt.Sprintf("Error while reading line #%d.\n\tError: %v", lineNumber, e)
			return nil, errors.New(msg)
		}

		strategy.rounds = append(strategy.rounds, round)
	}

	file.Close()
	return &strategy, nil
}

// Parse a round of the provided strategy
func parseRound(roundDescription []string) (Round, error) {
	var round Round = Round{}
	if len(roundDescription) != 2 {
		msg := fmt.Sprintf("A round is expected to have exactly two shapes, but %d were given: %v", len(roundDescription), roundDescription)
		return round, errors.New(msg)
	}

	abc, e := parseABC(roundDescription[0])
	if e != nil {
		return round, e
	}

	xyz, e := parseXYZ(roundDescription[1])
	if e != nil {
		return round, e
	}

	round.opponent = abc
	round.me = xyz
	return round, nil
}

// Parse 1-letter string as A, B or C
func parseABC(str string) (ABC, error) {
	switch str {
	case "A":
		return A, nil
	case "B":
		return B, nil
	case "C":
		return C, nil
	default:
		{
			var null ABC
			msg := fmt.Sprintf("Unexpected ABC description. Expected \"A\", \"B\" or \"C\", but received: \"%s\"", str)
			return null, errors.New(msg)
		}
	}
}

// Parse 1-letter string as X, Y or Z
func parseXYZ(str string) (XYZ, error) {
	switch str {
	case "X":
		return X, nil
	case "Y":
		return Y, nil
	case "Z":
		return Z, nil
	default:
		{
			var null XYZ
			msg := fmt.Sprintf("Unexpected XYZ description. Expected \"X\", \"Y\" or \"Z\", but received: \"%s\"", str)
			return null, errors.New(msg)
		}
	}
}
