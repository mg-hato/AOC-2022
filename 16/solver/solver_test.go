package solver

import (
	c "aoc/common"
	m "aoc/d16/models"
	ts "aoc/testers"
	"testing"
)

func test_valves() []m.Valve {
	return []m.Valve{
		{ID: "AA", Flow_rate: 0, Tunnels: []string{"BB", "CC"}},
		{ID: "BB", Flow_rate: 4, Tunnels: []string{"AA", "CC"}},
		{ID: "CC", Flow_rate: 0, Tunnels: []string{"AA", "BB", "DD"}},
		{ID: "DD", Flow_rate: 10, Tunnels: []string{"CC"}},
	}
}

func test_graph(deadline int) *graph {
	return make_graph(test_valves(), deadline)
}

func TestD16_FindAllDistances(t *testing.T) {
	expected_distances := map[c.Pair[string, string]]int{
		c.MakePair("AA", "BB"): 1,
		c.MakePair("AA", "CC"): 1,
		c.MakePair("BB", "CC"): 1,
		c.MakePair("CC", "DD"): 1,
		c.MakePair("AA", "DD"): 2,
		c.MakePair("BB", "DD"): 2,
	}

	// Make reverse pairs
	for key, value := range expected_distances {
		expected_distances[c.MakePair(key.Second, key.First)] = value
	}

	for _, node := range test_valves() {
		expected_distances[c.MakePair(node.ID, node.ID)] = 0
	}

	ts.AssertEqualWithEqFunc(
		t,
		findAllDistances(test_valves()),
		expected_distances,
		c.MapEqual[c.Pair[string, string], int],
	)
}

func TestD16_optimistic_pressure_release_approximation(t *testing.T) {
	g := test_graph(10)
	s := make_initial_state(test_valves(), 1)

	ts.AssertEqual(t, s.current_time, 1)
	ts.AssertEqualWithEqFunc(t, g.get_valve_greedy_opening_schedule(s), []string{"DD", "BB"}, c.ArrayEqual[string])

	// nearest non-zero valve is at distance 1 (namely BB)
	// reaching it and opening it would be done at minute 3
	// reaching next and opening it would be done at minute 5 at the earliest (optimistic estimation)
	// optimistically, we assume that we open valves with pressure 10 and 4 at minutes 3 and 5, respectively
	// until the deadline, that would release: (10+1 - 3) * 10 + (10+1 - 5) * 4 = 104 units of pressure
	ts.AssertEqual(t, g.approximate_total_optimistic_pressure_release(s), 104)

	g = test_graph(5) // now let's set deadline to be 5 minutes

	// same logic but deadline is now 5: (5+1 - 3) * 10 + (5+1 - 5) * 4 = 34 units of pressure
	ts.AssertEqual(t, g.approximate_total_optimistic_pressure_release(s), 34)

	g = test_graph(2) // now let's set deadling to be 2 minutes

	// with optimistic approximation, the nearest non-zero valve would be opened at minute 3, which is not enough time, so 0 pressure is optimistically approximated
	ts.AssertEqual(t, g.approximate_total_optimistic_pressure_release(s), 0)

	// now let us have deadline = 10 and modify the completion time of starting action to minute 6
	g = test_graph(10)
	s.agents[0].completion_time = 6

	// the nearest non-zero pressure would be reached & opened at minute 8
	// next one at minute 10 (optimistically)
	// hence the approximated pressure release would be (10+1 - 8) * 10 + (10+1 - 10) * 4 = 34
	ts.AssertEqual(t, g.approximate_total_optimistic_pressure_release(s), 34)

	// now let's assume that we already have pressure per minute releasing of 17 and that at the current minute (i.e. minute 1) 999 units of pressure are already released
	s.pressure_per_minute = 17
	s.pressure_released = 999

	// with these two in mind the optimistic pressure release get's increased by the "secured pressure release"
	// i.e. the amount that is guaranteed to be released by the deadline time based on already released pressure & pressure per minute release (and current agent's contribution which is zero in this case)
	// 999 + (10+1-1) * 17 = 1169
	ts.AssertEqual(t, g.get_secured_pressure_release(s), 1_169)
	ts.AssertEqual(t, g.approximate_total_optimistic_pressure_release(s), 34+g.get_secured_pressure_release(s))

	// lastly, move forward the current state to minute 6 (i.e. the time of the current's action completion)
	s = g.move_forward(s)
	ts.AssertEqual(t, s.current_time, 6)
	ts.AssertEqual(t, s.pressure_per_minute, 17)
	ts.AssertEqual(t, s.pressure_released, 1_084) // 999 + (6-1)*17 = 999 + 85 = 1084
}
