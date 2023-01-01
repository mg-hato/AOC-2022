package main

import (
	"aoc/reading"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type MotionSeriesReader struct {
	err         error
	data        []Motion
	move_regexp *regexp.Regexp
	line_number int
}

// CONSTRUCTOR

func NewMotionSeriesReader() reading.ReaderAoC2022[[]Motion] {
	return &MotionSeriesReader{
		err:         nil,
		data:        []Motion{},
		move_regexp: regexp.MustCompile("^([RLUD]) (\\d+)$"),
		line_number: 0,
	}
}

// PUBLIC METHODS

func (msr *MotionSeriesReader) Error() error {
	return msr.err
}

func (msr *MotionSeriesReader) PerformFinalValidation() error {
	return nil
}

func (msr *MotionSeriesReader) Done() bool {
	return msr.Error() != nil
}

func (msr *MotionSeriesReader) ProvideLine(line string) {
	msr.line_number++
	if move := msr.move_regexp.FindStringSubmatch(line); len(move) == 3 {
		steps, _ := strconv.Atoi(move[2])
		msr.data = append(msr.data, motion(move[1], steps))
	} else {
		msr.err = errors.New(fmt.Sprintf("Error: Could not interpret line #%d. Line: \"%s\"", msr.line_number, line))
	}

}

func (msr *MotionSeriesReader) FinishAndGetInputData() []Motion {
	return msr.data
}
