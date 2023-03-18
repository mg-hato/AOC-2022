package main

import . "aoc/functional"

func FindFirstSequenceOfDifferent(n int) func(string) (int, error) {
	return func(datastream string) (int, error) {

		// Process first n
		fc := FrequencyCounter{make(map[byte]int), 0}
		ForEach(func(b byte) { fc.AddElement(b) }, Take(n, []byte(datastream)))

		var i int = n
		for {
			// If frequency-counter detects n different elements, return i
			if fc.different == n {
				return i, nil
			}

			// If i still within length of datastream, continue processing, otherwise return -1
			if i < len(datastream) {
				fc.RemoveElement(datastream[i-n])
				fc.AddElement(datastream[i])
				i++
			} else {
				return -1, nil
			}
		}
	}
}
