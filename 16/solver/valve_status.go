package solver

type valve_status = int

const (
	closed valve_status = iota
	scheduled
	opened
	zero_pressure
)
