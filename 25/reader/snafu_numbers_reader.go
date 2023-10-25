package reader

import (
	"aoc/common"
	"aoc/d25/models"
	"aoc/reading"
	"fmt"
	"regexp"
)

type snafu_numbers_reader struct {
	err         error
	line_number int

	empty_re *regexp.Regexp
	snafu_re *regexp.Regexp

	numbers []string
}

func SnafuNumbersReader() reading.ReaderAoC2022[common.Envelope[[]string]] {
	return &snafu_numbers_reader{
		empty_re: regexp.MustCompile(`^ *$`),
		snafu_re: regexp.MustCompile(`^ *([012=-]+) *$`),
		numbers:  make([]string, 0),
	}
}

func (snr snafu_numbers_reader) Error() error {
	return snr.err
}

func (snr snafu_numbers_reader) PerformFinalValidation() error {
	return nil
}

func (snr snafu_numbers_reader) Done() bool {
	return snr.Error() != nil
}
func (snr *snafu_numbers_reader) ProvideLine(line string) {
	snr.line_number++
	switch {
	case snr.empty_re.MatchString(line):
	case snr.snafu_re.MatchString(line):
		snr.numbers = append(snr.numbers, snr.snafu_re.FindStringSubmatch(line)[1])
	default:
		snr.err = fmt.Errorf("bad line %d error", snr.line_number)
	}
}

func (snr snafu_numbers_reader) FinishAndGetInputData() common.Envelope[[]string] {
	return models.CreateSnafuNumbersEnvelope(snr.numbers)
}
