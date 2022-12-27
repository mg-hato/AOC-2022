package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadList(filename string) (List, error) {
	caloryList := List{calories: make([][]int, 0)}

	// Try to open the file
	file, e := os.Open(filename)
	if e != nil {
		return caloryList, e
	}

	defer file.Close()

	// Regular expression for the empty-line
	empty_line_regexp := regexp.MustCompile("^ *$")

	// Regular expression for a line that contains a non-negative number
	number_line_regexp := regexp.MustCompile("^\\d+$")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var line_number int = 0
	for scanner.Scan() {

		line_number++

		if empty_line_regexp.MatchString(scanner.Text()) {
			caloryList.separate()
		} else if number_line_regexp.MatchString(scanner.Text()) {
			item, _ := strconv.Atoi(scanner.Text())
			caloryList.addCaloryItem(item)
		} else {
			return caloryList, error_BadInputLine(line_number, scanner.Text())
		}
	}

	return caloryList, nil
}

func error_BadInputLine(line_number int, line string) error {
	message := fmt.Sprintf("Bad line of input on line %d. Line is: \"%s\"", line_number, line)
	return errors.New(message)
}

func (list *List) addCaloryItem(item int) {
	if len(list.calories) == 0 {
		list.calories = append(list.calories, []int{})
	}
	size := len(list.calories)
	list.calories[size-1] = append(list.calories[size-1], item)
}

func (list *List) separate() {
	if size := len(list.calories); size > 0 && len(list.calories[size-1]) > 0 {
		list.calories = append(list.calories, []int{})
	}
}
