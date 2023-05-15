package models

// A, B or C
type LeftSymbol int

const (
	A LeftSymbol = iota
	B
	C
)

func (left_symbol LeftSymbol) String() string {
	left_symbol_to_string := [...]string{"A", "B", "C"}
	if 0 <= int(left_symbol) && int(left_symbol) < len(left_symbol_to_string) {
		return left_symbol_to_string[int(left_symbol)]
	} else {
		return "UNKNOWN"
	}
}
