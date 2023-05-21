package reader

type reading_mode int

const (
	read_containers reading_mode = iota
	read_move_instructions
)
