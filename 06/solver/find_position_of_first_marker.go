package solver

import (
	f "aoc/functional"
	"errors"
)

// Position of the first marker is the first position where the `n` most recently received characters are different
func FindPositionOfTheFirstMarker(n int) func(string) (int, error) {
	return func(datastream string) (int, error) {

		fc := make_frequency_counter()
		// Process first n
		f.ForEach(fc.addElement, f.Take(n, []byte(datastream)))

		var i int = n
		for {
			if fc.different_count == n {
				return i, nil
			} else if i < len(datastream) {
				fc.removeElement(datastream[i-n])
				fc.addElement(datastream[i])
				i++
			} else {
				return -1, errors.New("solver error: no marker found")
			}
		}
	}
}
