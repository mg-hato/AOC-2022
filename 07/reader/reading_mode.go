package reader

type reading_mode = int

const (
	first_command reading_mode = iota
	new_command
	ls_command
)
