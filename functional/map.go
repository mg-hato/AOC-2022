package functional

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
