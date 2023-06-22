package solver

type noop struct{}

func (noop) get_next_actions(valve_statuses map[string]valve_status) []action {
	return []action{noop{}}
}

func (noop) get_valve_id() string {
	return ""
}
