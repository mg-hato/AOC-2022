package reader

import (
	c "aoc/common"
	m "aoc/d12/models"
	"aoc/reading"
	"regexp"
	"strings"
)

type terrain_reader struct {
	e       error
	terrain []string

	terrain_re  *regexp.Regexp
	empty_re    *regexp.Regexp
	line_number int
}

func TerrainReader() reading.ReaderAoC2022[m.SolverInput] {
	return &terrain_reader{
		terrain:    make([]string, 0),
		terrain_re: regexp.MustCompile("^([a-zSE]+) *$"),
		empty_re:   regexp.MustCompile("^ *$"),
	}
}

func (tr *terrain_reader) Error() error {
	return tr.e
}

func (tr *terrain_reader) PerformFinalValidation() error {

	// Check that some data is read
	if len(tr.terrain) == 0 {
		return no_terrain_was_given_error()
	}

	// Check length uniformity
	if !c.All(func(row string) bool { return len(row) == len(tr.terrain[0]) }, tr.terrain) {
		return terrain_map_is_not_rectangular_error()
	}

	// Count letter 'S'
	if s_count := c.Sum(c.Map(func(s string) int { return strings.Count(s, "S") }, tr.terrain)); s_count != 1 {
		return letter_expected_exactly_once_error('S', s_count)
	}

	// Count letter 'E'
	if e_count := c.Sum(c.Map(func(s string) int { return strings.Count(s, "E") }, tr.terrain)); e_count != 1 {
		return letter_expected_exactly_once_error('E', e_count)
	}

	return nil
}

func (tr *terrain_reader) Done() bool {
	return tr.Error() != nil
}

func (tr *terrain_reader) ProvideLine(line string) {
	tr.line_number++
	submatches := tr.terrain_re.FindStringSubmatch(line)
	switch {
	case submatches != nil:
		tr.terrain = append(tr.terrain, submatches[1])
	case tr.empty_re.MatchString(line):
	default:
		tr.e = invalid_line_error(tr.line_number, line)
	}
}

func (tr *terrain_reader) FinishAndGetInputData() m.SolverInput {
	return m.TerrainEnvelope(tr.terrain...)
}
