package reader

import (
	"aoc/reading"
	"regexp"
)

type jet_pattern_reader struct {
	err  error
	done bool

	line_number int

	jet_pattern string

	empty_re *regexp.Regexp

	jet_pattern_re *regexp.Regexp
}

func JetPatternReader() reading.ReaderAoC2022[string] {
	return &jet_pattern_reader{
		empty_re: regexp.MustCompile(`^ *$`),

		jet_pattern_re: regexp.MustCompile(`^ *((?:<|>)+) *$`),
	}
}

func (jpr jet_pattern_reader) Error() error {
	return jpr.err
}

func (jpr jet_pattern_reader) PerformFinalValidation() error {
	if !jpr.done {
		return jet_pattern_not_read_validation_error()
	}
	return nil
}

func (jpr jet_pattern_reader) Done() bool {
	return jpr.done || jpr.Error() != nil
}

func (jpr *jet_pattern_reader) ProvideLine(line string) {
	jpr.line_number++

	switch {
	case jpr.empty_re.MatchString(line):
	case jpr.jet_pattern_re.MatchString(line):
		jpr.jet_pattern = jpr.jet_pattern_re.FindStringSubmatch(line)[1]
		jpr.done = true
	default:
		jpr.err = bad_line_reader_error(jpr.line_number, line)
	}
}

func (jpr *jet_pattern_reader) FinishAndGetInputData() string {
	return jpr.jet_pattern
}
