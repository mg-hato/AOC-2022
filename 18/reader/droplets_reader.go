package reader

import (
	c "aoc/common"
	m "aoc/d18/models"
	"aoc/reading"
	"regexp"
	"strconv"
)

type droplets_reader struct {
	err error

	line_number int

	empty_re   *regexp.Regexp
	droplet_re *regexp.Regexp

	droplets []m.Droplet
}

func DropletsReader() reading.ReaderAoC2022[m.SolverInput] {
	return &droplets_reader{
		empty_re:   regexp.MustCompile(`^ *$`),
		droplet_re: regexp.MustCompile(`^ *(\d+) *, *(\d+) *, *(\d+) *$`),
	}
}

func (dr droplets_reader) Error() error {
	return dr.err
}

func (dr droplets_reader) PerformFinalValidation() error {
	return nil
}

func (dr droplets_reader) Done() bool {
	return dr.Error() != nil
}

func (dr *droplets_reader) ProvideLine(line string) {
	dr.line_number++

	switch {
	case dr.empty_re.MatchString(line):
	case dr.droplet_re.MatchString(line):
		coordinates := c.Map(
			func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			},
			dr.droplet_re.FindStringSubmatch(line)[1:],
		)
		dr.droplets = append(dr.droplets, m.MakeDroplet(coordinates[0], coordinates[1], coordinates[2]))
	default:
		dr.err = bad_line(dr.line_number, line)
	}
}

func (dr droplets_reader) FinishAndGetInputData() m.SolverInput {
	return m.DropletsEnvelope(dr.droplets...)
}
