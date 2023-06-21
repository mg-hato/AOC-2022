package reader

import (
	c "aoc/common"
	m "aoc/d14/models"
	"aoc/reading"
	"regexp"
	"strconv"
)

func RockStructureReader() reading.ReaderAoC2022[m.SolverInput] {
	return &rock_structure_reader{

		sand_source: c.MakePair(500, 0),

		empty_re:     regexp.MustCompile("^ *$"),
		rock_path_re: regexp.MustCompile(`^ *\d+ *, *\d+(?: *-> *\d+ *, *\d+)+$`),

		rock_path_coordinates_re: regexp.MustCompile(`(\d+) *, *(\d+)`),
	}
}

type rock_structure_reader struct {
	err error

	line_number int

	sand_source c.Pair[int, int]

	rock_structures []m.RockStructure

	empty_re     *regexp.Regexp
	rock_path_re *regexp.Regexp

	rock_path_coordinates_re *regexp.Regexp
}

func (rsr rock_structure_reader) Error() error {
	return rsr.err
}

func (rsr rock_structure_reader) PerformFinalValidation() error {
	for _, validation_function := range []func([]m.RockStructure) error{
		verify_that_rock_paths_are_horizontal_vertical,
		verify_that_no_rock_formation_overlaps_sand_source(rsr.sand_source),
	} {
		if validation_error := validation_function(rsr.rock_structures); validation_error != nil {
			return validation_error
		}
	}
	return nil
}

func (rsr rock_structure_reader) Done() bool {
	return rsr.Error() != nil
}

func (rsr *rock_structure_reader) ProvideLine(line string) {
	rsr.line_number++

	switch {
	case rsr.empty_re.MatchString(line):
	case rsr.rock_path_re.MatchString(line):
		rock_structure := c.Map(
			func(submatches []string) c.Pair[int, int] {
				first, _ := strconv.Atoi(submatches[1])
				second, _ := strconv.Atoi(submatches[2])
				return c.MakePair(first, second)
			},
			rsr.rock_path_coordinates_re.FindAllStringSubmatch(line, -1),
		)
		rsr.rock_structures = append(rsr.rock_structures, rock_structure)
	default:
		rsr.err = bad_line_reading_error(rsr.line_number, line)
	}
}

func (rsr rock_structure_reader) FinishAndGetInputData() m.SolverInput {
	return m.RockStructureEnvelope(rsr.rock_structures)
}
