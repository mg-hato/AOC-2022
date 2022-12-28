package main

type FrequencyCounter struct {
	frequencies map[byte]int
	different   int
}

func (fc *FrequencyCounter) AddElement(b byte) {
	fc.frequencies[b]++
	if fc.frequencies[b] == 1 {
		fc.different++
	}
}

func (fc *FrequencyCounter) RemoveElement(b byte) {
	fc.frequencies[b]--
	if fc.frequencies[b] == 0 {
		fc.different--
	} else if fc.frequencies[b] == -1 {
		fc.frequencies[b] = 0
	}
}
