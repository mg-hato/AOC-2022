package models

type repeating_jet_pattern struct {
	jet_pattern string
	current     int
}

func RepeatingJetPattern(jet_pattern string) *repeating_jet_pattern {
	return &repeating_jet_pattern{jet_pattern: jet_pattern}
}

func (rjp *repeating_jet_pattern) Next() (func(Shape) Shape, int) {
	var shape_transition_function func(Shape) Shape

	jet_index := rjp.current

	rjp.current++
	if rjp.current == len(rjp.jet_pattern) {
		rjp.current = 0
	}

	switch rjp.jet_pattern[jet_index] {
	case '<':
		shape_transition_function = MoveLeft
	case '>':
		shape_transition_function = MoveRight
	}

	return shape_transition_function, jet_index
}
