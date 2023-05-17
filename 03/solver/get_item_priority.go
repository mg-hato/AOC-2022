package solver

func get_item_priority(item rune) int {
	if 'a' <= item && item <= 'z' {
		return int(item) - int('a') + 1
	} else if 'A' <= item && item <= 'Z' {
		return int(item) - int('A') + 27
	}
	return 0
}
