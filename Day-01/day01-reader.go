package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadCaloryList(filename string) (*CaloryList, error) {
	caloryList := CaloryList{list: make([][]int, 0)}

	// Try to open the file
	file, e := os.Open(filename)
	if e != nil {
		return nil, e
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lineNumber int = 0
	for scanner.Scan() {

		// If error during reading, return
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		lineNumber++
		line := scanner.Text()

		// On empty line, separate; on non-empty parse the caloric value
		switch line {
		case "":
			caloryList.separate()
		default:
			{
				if err := caloryList.addCaloryItem(line); err != nil {
					return nil, errors.New(getErrMsg(lineNumber, line, err))
				}
			}
		}
	}

	return &caloryList, nil
}

func getErrMsg(lineNumber int, line string, e error) string {
	return fmt.Sprintf("Error occurred reading/parsing line #%d. Line is: \"%s\"\nError is: %v", lineNumber, line, e)
}

func (caloryList *CaloryList) addCaloryItem(line string) error {
	item, err := strconv.Atoi(line)
	if err != nil {
		return err
	}

	if len(caloryList.list) == 0 {
		caloryList.list = append(caloryList.list, []int{})
	}
	size := len(caloryList.list)
	caloryList.list[size-1] = append(caloryList.list[size-1], item)

	return nil
}

func (caloryList *CaloryList) separate() {
	if size := len(caloryList.list); size > 0 && len(caloryList.list[size-1]) > 0 {
		caloryList.list = append(caloryList.list, []int{})
	}
}
