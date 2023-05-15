package reader

import (
	m "aoc/day01/models"
	e "aoc/envelope"
	"aoc/reading"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type calory_list_reader struct {
	err         error
	calory_list m.CaloryList
	current_idx int
	line_number int

	empty_line_re *regexp.Regexp
	calory_re     *regexp.Regexp
}

// Constructor function for calory list reader
func CaloryListReader() reading.ReaderAoC2022[e.Envelope[m.CaloryList]] {
	return &calory_list_reader{
		err: nil,

		calory_list: [][]int{},
		current_idx: 0,
		line_number: 0,

		empty_line_re: regexp.MustCompile("^ *$"),
		calory_re:     regexp.MustCompile(`^ *([1-9]\d*|0) *$`),
	}
}

func (clr calory_list_reader) Error() error {
	return clr.err
}

func (clr calory_list_reader) PerformFinalValidation() error {
	if len(clr.calory_list) == 0 {
		return errors.New("Error: calory list is empty")
	}
	return clr.Error()
}

func (clr calory_list_reader) Done() bool {
	return false
}

func (clr *calory_list_reader) ProvideLine(line string) {
	clr.line_number++
	if clr.empty_line_re.MatchString(line) {
		clr.current_idx = len(clr.calory_list)
	} else if submatches := clr.calory_re.FindStringSubmatch(line); len(submatches) == 2 {
		calory, _ := strconv.Atoi(submatches[1])
		clr.addCalory(calory)
	} else {
		clr.err = fmt.Errorf(
			"Error: could not interpret line #%d: %s",
			clr.line_number, line,
		)
	}
}

func (clr *calory_list_reader) addCalory(calory int) {
	if len(clr.calory_list) == clr.current_idx {
		clr.calory_list = append(clr.calory_list, []int{calory})
	} else {
		clr.calory_list[clr.current_idx] = append(clr.calory_list[clr.current_idx], calory)
	}
}

func (clr calory_list_reader) FinishAndGetInputData() e.Envelope[m.CaloryList] {
	return m.CreateCaloryListEnvelope(clr.calory_list)
}
