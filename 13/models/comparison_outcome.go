package models

type comparison_outcome = int

const (
	in_order comparison_outcome = iota
	same
	not_in_order
)
