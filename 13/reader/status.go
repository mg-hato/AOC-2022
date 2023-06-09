package reader

type status = int

const (
	first_packet status = iota
	second_packet
	empty_line
)
