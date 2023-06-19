package reader

import (
	m "aoc/d15/models"
	e "aoc/envelope"
	"aoc/functional"
	"aoc/reading"
	"regexp"
	"strconv"
)

type sensor_reports_reader struct {
	err error

	line_number int

	empty_re  *regexp.Regexp
	report_re *regexp.Regexp

	reports []m.SensorReport
}

func SensorReportsReader() reading.ReaderAoC2022[e.Envelope[[]m.SensorReport]] {
	return &sensor_reports_reader{
		empty_re:  regexp.MustCompile(`^ *$`),
		report_re: regexp.MustCompile(`^ *Sensor at x=(-?\d+), *y=(-?\d+) *: *closest beacon is at x=(-?\d+), *y=(-?\d+) *$`),
	}
}

func (srr sensor_reports_reader) Error() error {
	return srr.err
}

func (srr sensor_reports_reader) PerformFinalValidation() error {
	for _, validation_func := range []func([]m.SensorReport) error{
		verify_that_no_two_sensors_overlap,
		verify_beacons_are_the_closest,
		verify_no_sensor_has_multiple_equidistant_beacons,
	} {
		if err := validation_func(srr.reports); err != nil {
			return err
		}
	}
	return nil
}

func (srr sensor_reports_reader) Done() bool {
	return srr.Error() != nil
}

func (srr *sensor_reports_reader) ProvideLine(line string) {
	srr.line_number++

	switch {
	case srr.empty_re.MatchString(line):
	case srr.report_re.MatchString(line):
		numbers := functional.Map(
			func(s string) int { i, _ := strconv.Atoi(s); return i },
			srr.report_re.FindStringSubmatch(line)[1:],
		)
		report := m.SensorReport{
			Sensor: m.MakePoint(numbers[0], numbers[1]),
			Beacon: m.MakePoint(numbers[2], numbers[3]),
		}
		srr.reports = append(srr.reports, report)
	default:
		srr.err = bad_line_reader_error(line, srr.line_number)
	}
}

func (srr sensor_reports_reader) FinishAndGetInputData() e.Envelope[[]m.SensorReport] {
	return m.SensorReportsEnvelope(srr.reports...)
}
