package functional

// Get values of the map in an array
func GetValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	var i int
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Get keys of the map in an array
func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// Checks if the two maps have the same key-value mappings
func MapEqual[K, V comparable](lhs, rhs map[K]V) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k, v := range lhs {
		if rv, ok := rhs[k]; !ok || v != rv {
			return false
		}
	}

	return true
}
