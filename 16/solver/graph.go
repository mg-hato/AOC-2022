package solver

import (
	c "aoc/common"
	m "aoc/d16/models"
	"sort"
)

type graph struct {
	valves    map[string]m.Valve
	distances map[c.Pair[string, string]]int
	deadline  int
}

func make_graph(valves []m.Valve, deadline int) *graph {
	return &graph{
		valves:    c.AssociateBy(valves, func(v m.Valve) string { return v.ID }),
		distances: findAllDistances(valves),
		deadline:  deadline,
	}
}

func (g *graph) get_distance(from, to string) int {
	return g.distances[c.MakePair(from, to)]
}

func (g *graph) expand_state(S state) []state {
	if S.current_time > g.deadline {
		return []state{}
	}
	expanded_states := []state{S}
	for _, A := range S.agents {
		expanded_states = c.FlatMap(func(st state) []state { return g.expand_state_on_agent(st, A) }, expanded_states)
	}

	return c.Map(func(st state) state {
		st = g.move_forward(st)
		st.oprp = g.approximate_total_optimistic_pressure_release(st)
		return st
	}, expanded_states)
}

func (g *graph) move_forward(S state) state {
	updated_time := c.Minimum(c.Map(func(A agent) int { return A.completion_time }, S.agents))
	updated_time = c.Min(updated_time, g.deadline+1)
	S.pressure_released += c.Max(updated_time-S.current_time, 0) * S.pressure_per_minute
	S.current_time = updated_time
	c.ForEach(func(A agent) {
		_, is_open_valve_action := A.act.(open_valve)
		if A.completion_time == S.current_time && is_open_valve_action {
			valve_id := A.act.get_valve_id()
			S.valve_statuses[valve_id] = opened
			S.pressure_per_minute += g.valves[valve_id].Flow_rate
		}
	}, S.agents)
	return S
}

func (g *graph) expand_state_on_agent(S state, A agent) []state {

	// This agent's action is not completed, so we do not expand it yet
	if S.current_time != A.completion_time {
		return []state{S}
	}

	return c.Map(func(act action) state {
		new_state := S.make_copy()

		var completion_time int

		switch act := act.(type) {
		case open_valve:
			completion_time = A.completion_time + g.get_distance(A.act.get_valve_id(), act.id) + 1
			new_state.valve_statuses[act.id] = scheduled
		default:
			completion_time = g.deadline + 1
		}

		new_state.agents[A.agent_id] = agent{
			agent_id:        A.agent_id,
			act:             act,
			completion_time: completion_time,
		}

		return new_state
	}, A.act.get_next_actions(S.valve_statuses))
}

// This represents the pressure that will be released by the deadline
// as a consequence of agents' current actions
func (g *graph) get_agents_current_action_contribution(S state) int {
	return c.Sum(c.Map(func(A agent) int {
		vid := A.act.get_valve_id()
		if status, ok := S.valve_statuses[vid]; ok && status == scheduled {
			return g.valves[vid].Flow_rate * c.Max(g.deadline+1-A.completion_time, 0)
		}
		return 0
	}, S.agents))
}

func (g *graph) get_secured_pressure_release(S state) int {
	return S.pressure_released +
		S.pressure_per_minute*c.Max(g.deadline+1-S.current_time, 0) +
		g.get_agents_current_action_contribution(S)
}

func (g *graph) get_valve_greedy_opening_schedule(S state) []string {
	schedule := c.Filter(
		func(valve_id string) bool { return S.valve_statuses[valve_id] == closed },
		c.GetKeys(S.valve_statuses),
	)
	sort.Slice(schedule, func(i, j int) bool {
		return g.valves[schedule[i]].Flow_rate > g.valves[schedule[j]].Flow_rate
	})

	return schedule
}

func (g *graph) approximate_total_optimistic_pressure_release(S state) int {
	secured_pressure_release := g.get_secured_pressure_release(S)

	schedule := g.get_valve_greedy_opening_schedule(S)
	active_agents := S.get_active_agents()

	if len(schedule) == 0 || len(active_agents) == 0 {
		return secured_pressure_release
	}

	agents_availabilities := c.Map(func(a agent) int {
		return c.Minimum(c.Map(
			func(valve_id string) int {
				return a.completion_time + g.get_distance(a.act.get_valve_id(), valve_id) + 1
			},
			schedule,
		))
	}, active_agents)

	AOPRP := 0 // aggregated optimistic pressure release prediction
	index, time := 0, S.current_time
	for index < len(schedule) && time <= g.deadline {
		for i := 0; i < len(agents_availabilities); i++ {
			if agents_availabilities[i] == time && index < len(schedule) {
				AOPRP += g.valves[schedule[index]].Flow_rate * c.Max(g.deadline+1-time, 0)
				agents_availabilities[i] += 2
				index++
			}
		}
		time++
	}

	return secured_pressure_release + AOPRP
}
