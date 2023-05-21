package solver

type frequency_counter struct {
	frequency_map   map[byte]int
	different_count int
}

func make_frequency_counter() *frequency_counter {
	return &frequency_counter{
		frequency_map:   make(map[byte]int),
		different_count: 0,
	}
}

func (fc *frequency_counter) addElement(element byte) {
	if fc.frequency_map[element] == 0 {
		fc.different_count++
	}
	fc.frequency_map[element]++
}

func (fc *frequency_counter) removeElement(element byte) {
	fc.frequency_map[element]--
	if fc.frequency_map[element] == 0 {
		fc.different_count--
	}
}
