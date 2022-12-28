package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func ReadStrategy(filename string) (EncryptedStrategyGuide, error) {
	strategy := EncryptedStrategyGuide{rounds: []Round{}}
	file, e := os.Open(filename)

	// Try to open the file
	if e != nil {
		return strategy, e
	}

	// A line is described with opponent's move (A, B or C) and my move (X, Y or Z)
	round_regexp := regexp.MustCompile("^([ABC]) ([XYZ])$")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var line_number int
	for scanner.Scan() {
		line_number++

		if !round_regexp.MatchString(scanner.Text()) {
			return strategy, error_BadInputLine(scanner.Text(), line_number)
		}

		round := parseRound(round_regexp.FindStringSubmatch(scanner.Text())[1:])
		strategy.rounds = append(strategy.rounds, round)
	}

	return strategy, nil
}

// Parse one round of the given strategy guide
func parseRound(round_as_strings []string) Round {
	return Round{parseABC(round_as_strings[0]), parseXYZ(round_as_strings[1])}
}

// Parse 1-letter string as A, B or C
func parseABC(str string) ABC {
	switch str {
	case "A":
		return A
	case "B":
		return B
	default:
		return C
	}
}

// Parse 1-letter string as X, Y or Z
func parseXYZ(str string) XYZ {
	switch str {
	case "X":
		return X
	case "Y":
		return Y
	default:
		return Z
	}
}

func error_BadInputLine(line string, line_number int) error {
	message := fmt.Sprintf("Bad line of input describing one round on line %d. Line is: \"%s\"", line_number, line)
	return errors.New(message)
}
