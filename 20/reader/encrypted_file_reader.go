package reader

import (
	"aoc/common"
	m "aoc/d20/models"
	"aoc/reading"
	"regexp"
	"strconv"
)

type encrypted_file_reader struct {
	err         error
	line_number int

	empty_re  *regexp.Regexp
	number_re *regexp.Regexp

	encrypted_file m.EncryptedFile
}

func EncryptedFileReader() reading.ReaderAoC2022[m.SolverInput] {
	return &encrypted_file_reader{
		empty_re:  regexp.MustCompile("^ *$"),
		number_re: regexp.MustCompile(`^ *(-?\d+) *$`),
	}
}

func (efr encrypted_file_reader) Error() error {
	return efr.err
}

func (efr encrypted_file_reader) PerformFinalValidation() error {

	// Ensure there is exactly one 0
	if zero_count := common.Count(efr.encrypted_file, func(value int) bool { return value == 0 }); zero_count != 1 {
		return invalid_zero_count_validation_error(zero_count)
	}
	return nil
}

func (efr encrypted_file_reader) Done() bool {
	return efr.Error() != nil
}

func (efr *encrypted_file_reader) ProvideLine(line string) {
	efr.line_number++

	switch {
	case efr.empty_re.MatchString(line):
	case efr.number_re.MatchString(line):
		number, _ := strconv.Atoi(efr.number_re.FindStringSubmatch(line)[1])
		efr.encrypted_file = append(efr.encrypted_file, number)
	default:
		efr.err = bad_line_reader_error(efr.line_number, line)
	}
}

func (efr encrypted_file_reader) FinishAndGetInputData() m.SolverInput {
	return m.EncryptedFileEnvelope(efr.encrypted_file)
}
