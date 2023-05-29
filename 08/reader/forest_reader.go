package reader

import (
	m "aoc/day08/models"
	e "aoc/envelope"
	f "aoc/functional"
	"aoc/reading"
	"regexp"
)

type forest_reader struct {
	err    error
	forest m.Forest

	line_number   int
	forest_row_re *regexp.Regexp
}

func ForestReader() reading.ReaderAoC2022[e.Envelope[m.Forest]] {
	return &forest_reader{
		forest:        make(m.Forest, 0),
		forest_row_re: regexp.MustCompile(`^(\d+) *$`),
	}
}

func (fr forest_reader) Error() error {
	return fr.err
}

func (fr forest_reader) PerformFinalValidation() error {

	if len(fr.forest) == 0 {
		return empty_forest_error()
	}

	if row_lengths_are_equal := f.ArrayEqual(
		f.Map(func(row []byte) int { return len(row) }, fr.forest),
		f.Repeat(len(fr.forest[0]), len(fr.forest)),
	); !row_lengths_are_equal {
		return different_forest_row_lengths_error(fr.forest)
	}

	return nil
}

func (fr *forest_reader) Done() bool {
	return fr.err != nil
}

func (fr *forest_reader) ProvideLine(line string) {
	fr.line_number++
	submatches := fr.forest_row_re.FindStringSubmatch(line)
	if submatches == nil {
		fr.err = bad_line_error(fr.line_number, line)
		return
	}
	fr.forest = append(fr.forest, f.Map(func(b byte) byte { return b - '0' }, []byte(submatches[1])))
}

func (fr forest_reader) FinishAndGetInputData() e.Envelope[m.Forest] {
	return m.ForestEnvelope(fr.forest)
}
