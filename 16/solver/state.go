package solver

import (
	c "aoc/common"
	m "aoc/d16/models"
)

type state struct {
	agents         []agent
	valve_statuses map[string]valve_status

	current_time int

	pressure_released   int
	pressure_per_minute int

	oprp int // Optimistic pressure release prediction
}

func make_initial_state(valves []m.Valve, number_of_agents int) state {
	return state{
		valve_statuses: c.CreateKeyValueMap(
			valves,
			func(valve m.Valve) string { return valve.ID },
			func(valve m.Valve) valve_status {
				if valve.Flow_rate == 0 {
					return zero_pressure
				} else {
					return closed
				}
			},
		),

		agents: c.Map(
			func(agent_id int) agent {
				return agent{
					agent_id: agent_id,
					act:      start_action{},

					completion_time: 1,
				}
			},
			c.Range(0, number_of_agents),
		),

		current_time:        1,
		pressure_released:   0,
		pressure_per_minute: 0,

		oprp: 1,
	}
}

func (S state) make_copy() state {
	S.agents = c.ShallowCopy(S.agents)
	S.valve_statuses = c.AssociateWith(
		c.GetKeys(S.valve_statuses),
		func(id string) valve_status { return S.valve_statuses[id] },
	)
	return S
}

func (S state) get_active_agents() []agent {
	return c.Filter(
		func(ag agent) bool {
			_, is_noop := ag.act.(noop)
			return !is_noop
		},
		S.agents,
	)
}
