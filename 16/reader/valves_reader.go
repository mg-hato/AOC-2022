package reader

import (
	m "aoc/d16/models"
	"aoc/reading"
	"regexp"
	"strconv"
)

type valves_reader struct {
	err error

	line_number int

	valves []m.Valve

	empty_re *regexp.Regexp
	valve_re *regexp.Regexp

	id_xre        *regexp.Regexp
	flow_rate_xre *regexp.Regexp
}

func ValvesReader() reading.ReaderAoC2022[m.SolverInput] {
	return &valves_reader{
		empty_re: regexp.MustCompile(`^ *$`),
		valve_re: regexp.MustCompile(`^Valve [A-Z]{2} has flow rate=\d+; tunnels? leads? to valves? [A-Z]{2}(?: *, *[A-Z]{2})* *$`),

		id_xre:        regexp.MustCompile(`[A-Z]{2}`),
		flow_rate_xre: regexp.MustCompile(`flow rate=(\d+)`),
	}
}

func (vr valves_reader) Error() error {
	return vr.err
}

func (vr valves_reader) PerformFinalValidation() error {

	for _, validation_func := range []func([]m.Valve) error{
		verify_starting_valve_exists,
		verify_that_valves_have_unique_names,
		verify_that_no_valve_tunnel_leads_to_itself,
		verify_that_all_valve_tunnels_lead_to_distinct_valve,
		verify_that_valve_tunnels_lead_to_defined_valves,
	} {
		if validation_err := validation_func(vr.valves); validation_err != nil {
			return validation_err
		}
	}

	return nil
}

func (vr valves_reader) Done() bool {
	return vr.Error() != nil
}

func (vr *valves_reader) ProvideLine(line string) {
	vr.line_number++

	switch {
	case vr.empty_re.MatchString(line):
	case vr.valve_re.MatchString(line):
		vr.processValveLine(line)
	default:
		vr.err = bad_line_error(vr.line_number, line)
	}
}

func (vr *valves_reader) processValveLine(line string) {
	valve := m.Valve{}

	flow_rate, _ := strconv.Atoi(vr.flow_rate_xre.FindStringSubmatch(line)[1])
	valve.Flow_rate = flow_rate

	ids := vr.id_xre.FindAllString(line, -1)
	valve.ID = ids[0]
	valve.Tunnels = ids[1:]

	vr.valves = append(vr.valves, valve)
}

func (vr valves_reader) FinishAndGetInputData() m.SolverInput {
	return m.ValveEnvelope(vr.valves...)
}
