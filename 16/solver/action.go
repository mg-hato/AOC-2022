package solver

type action interface {
	get_next_actions(map[string]valve_status) []action
	get_valve_id() string
}
