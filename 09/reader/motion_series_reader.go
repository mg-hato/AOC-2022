package reader

import (
	m "aoc/d09/models"
	"aoc/reading"
	"fmt"
	"regexp"
	"strconv"
)

type motion_series_reader struct {
	err           error
	motion_series m.MotionSeries
	move_re       *regexp.Regexp
	empty_re      *regexp.Regexp
	line_number   int
}

func MotionSeriesReader() reading.ReaderAoC2022[m.SolverInput] {
	return &motion_series_reader{
		motion_series: make(m.MotionSeries, 0),
		move_re:       regexp.MustCompile(`^ *([RLUD]) +([1-9]\d*) *$`),
		empty_re:      regexp.MustCompile(`^ *$`),
	}
}

func (msr *motion_series_reader) Error() error {
	return msr.err
}

func (msr *motion_series_reader) PerformFinalValidation() error {
	return nil
}

func (msr *motion_series_reader) Done() bool {
	return msr.Error() != nil
}

func (msr *motion_series_reader) ProvideLine(line string) {
	msr.line_number++

	if msr.empty_re.MatchString(line) {
		return
	}

	submatches := msr.move_re.FindStringSubmatch(line)
	if submatches == nil {
		msr.err = fmt.Errorf(`error while reading motion series on line #%d: could not interpret line "%s"`, msr.line_number, line)
		return
	}

	direction := m.Direction(submatches[1])
	steps, _ := strconv.Atoi(submatches[2])

	msr.motion_series = append(msr.motion_series, m.MakeMotion(steps, direction))
}

func (msr *motion_series_reader) FinishAndGetInputData() m.SolverInput {
	return m.MotionSeriesEnvelope(msr.motion_series)
}
