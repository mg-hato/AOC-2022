package reader

import (
	c "aoc/common"
	m "aoc/d19/models"
	"aoc/reading"
	"regexp"
	"strconv"
)

type blueprints_reader struct {
	err error

	line_number int

	empty_re *regexp.Regexp

	blueprint_re *regexp.Regexp

	blueprints []m.Blueprint
}

func BlueprintsReader() reading.ReaderAoC2022[m.SolverInput] {
	return &blueprints_reader{
		empty_re: regexp.MustCompile(`^ *$`),

		blueprint_re: regexp.MustCompile(`^ *Blueprint (\d+): Each ore robot costs (\d+) ore\. Each clay robot costs (\d+) ore\. Each obsidian robot costs (\d+) ore and (\d+) clay\. Each geode robot costs (\d+) ore and (\d+) obsidian\. *$`),
	}
}

func (br blueprints_reader) Error() error {
	return br.err
}

func (br blueprints_reader) PerformFinalValidation() error {
	blueprint_ids := c.Map(func(b m.Blueprint) int { return b.ID }, br.blueprints)
	if !c.ArrayEqual(blueprint_ids, c.RangeInclusive(1, len(blueprint_ids))) {
		return invalid_blueprint_ids_validation_error(len(blueprint_ids))
	}
	return nil
}

func (br blueprints_reader) Done() bool {
	return br.Error() != nil
}

func (br *blueprints_reader) ProvideLine(line string) {
	br.line_number++

	switch {
	case br.empty_re.MatchString(line):
	case br.blueprint_re.MatchString(line):
		numbers := c.Map(
			func(s string) int { i, _ := strconv.Atoi(s); return i },
			br.blueprint_re.FindStringSubmatch(line)[1:],
		)
		br.blueprints = append(br.blueprints, m.MakeBlueprint(
			numbers[0],
			numbers[1],
			numbers[2],
			numbers[3],
			numbers[4],
			numbers[5],
			numbers[6],
		))

	default:
		br.err = bad_line_reader_error(br.line_number, line)
	}
}

func (br blueprints_reader) FinishAndGetInputData() m.SolverInput {
	return m.BlueprintsEnvelope(br.blueprints...)
}
