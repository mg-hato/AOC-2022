package solver

import c "aoc/common"

type start_action struct{}

func (start_action) get_next_actions(valve_statuses map[string]valve_status) []action {
	return append(
		c.Map(
			func(valve_id string) action { return open_valve{valve_id} },
			c.Filter(func(valve_id string) bool { return valve_statuses[valve_id] == closed }, c.GetKeys(valve_statuses)),
		),
		noop{},
	)
}

func (start_action) get_valve_id() string {
	return "AA"
}
