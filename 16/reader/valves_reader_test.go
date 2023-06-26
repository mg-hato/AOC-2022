package reader

import (
	m "aoc/d16/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD16_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(ValvesReader)).
		ProvideEqualityFunction(m.ValvesEnvelopeEqualityFunc).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.ValveEnvelope([]m.Valve{
			{ID: "AA", Flow_rate: 0, Tunnels: []string{"DD", "BB"}},
			{ID: "BB", Flow_rate: 13, Tunnels: []string{"CC", "AA"}},
			{ID: "CC", Flow_rate: 2, Tunnels: []string{"DD", "BB"}},
			{ID: "DD", Flow_rate: 20, Tunnels: []string{"CC", "AA", "EE"}},
			{ID: "EE", Flow_rate: 3, Tunnels: []string{"DD"}},
		}...))).
		AddTestCase("./tests/no_starting_valve.txt", ts.ExpectError[m.SolverInput](
			validation_error_prefix(),
			`no valve with ID "AA"`,
		)).
		AddTestCase("./tests/valve_defined_twice.txt", ts.ExpectError[m.SolverInput](
			validation_error_prefix(),
			"valve with name BB",
			"defined twice",
			"valve #2",
			"valve #4",
		)).
		AddTestCase("./tests/valve_has_self_loop.txt", ts.ExpectError[m.SolverInput](
			validation_error_prefix(),
			"valve EE",
			"leads to itself",
		)).
		AddTestCase("./tests/multiple_tunnels_leading_to_the_same_valve.txt", ts.ExpectError[m.SolverInput](
			validation_error_prefix(),
			"valve DD",
			"more than one tunnel leading to valve CC",
		)).
		AddTestCase("./tests/tunnel_to_undefined_valve.txt", ts.ExpectError[m.SolverInput](
			validation_error_prefix(),
			"valve AA",
			"tunnel",
			"undefined valve ZZ",
		)).
		RunReaderTests()
}
