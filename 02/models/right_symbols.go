package models

// X, Y or Z
type RightSymbol int

const (
	X RightSymbol = iota
	Y
	Z
)

func (right_symbol RightSymbol) String() string {
	right_symbol_to_string := [...]string{"X", "Y", "Z"}
	if 0 <= int(right_symbol) && int(right_symbol) < len(right_symbol_to_string) {
		return right_symbol_to_string[int(right_symbol)]
	} else {
		return "UNKNOWN"
	}
}
