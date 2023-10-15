package reader

import (
	"aoc/common"
	"aoc/d24/models"
	"aoc/reading"
	"fmt"
	"regexp"
)

type valley_map_reader struct {
	err         error
	line_number int

	done bool

	valley_map [][]rune

	empty_re, valley_row_re *regexp.Regexp
}

func ValleyMapReader() reading.ReaderAoC2022[common.Envelope[[][]rune]] {
	return &valley_map_reader{
		empty_re:      regexp.MustCompile(`^ *$`),
		valley_row_re: regexp.MustCompile(`^([#\.<>^v]+) *$`),
		valley_map:    make([][]rune, 0),
	}
}

func (vmr valley_map_reader) Error() error {
	return vmr.err
}
func (vmr valley_map_reader) PerformFinalValidation() error {
	for _, validation_func := range []func([][]rune) error{
		validate_map_is_not_empty,
		validate_square_map,
		validate_sides_and_exits,
	} {
		if validation_err := validation_func(vmr.valley_map); validation_err != nil {
			return validation_err
		}
	}
	return nil
}

func (vmr valley_map_reader) Done() bool {
	return vmr.Error() != nil || vmr.done
}

func (vmr *valley_map_reader) ProvideLine(line string) {
	vmr.line_number++
	switch {
	case vmr.empty_re.MatchString(line):
		if len(vmr.valley_map) > 0 {
			vmr.done = true
		}
	case vmr.valley_row_re.MatchString(line):
		valley_map_row := vmr.valley_row_re.FindStringSubmatch(line)[1]
		vmr.valley_map = append(vmr.valley_map, []rune(valley_map_row))
	default:
		vmr.err = fmt.Errorf(`bad line #%d: "%s"`, vmr.line_number, line)
	}
}

func (vmr valley_map_reader) FinishAndGetInputData() common.Envelope[[][]rune] {
	return models.CreateValleyMapEnvelope(vmr.valley_map)
}
