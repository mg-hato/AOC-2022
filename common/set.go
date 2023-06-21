package common

// Create a set from the array `arr`
// such that for every element `e` of `arr` the corresponding set element would be `keyf(e)`
func CreateSet[T any, K comparable](arr []T, keyf func(T) K) map[K]bool {
	return CreateKeyValueMap(arr, keyf, func(T) bool { return true })
}

func SetContains[K comparable](key K) func(map[K]bool) bool {
	return func(set map[K]bool) bool {
		return set[key]
	}
}

func SetEqual[K comparable](lhs, rhs map[K]bool) bool {
	return ArrayEqualInAnyOrder(
		Filter(func(key K) bool { return lhs[key] }, GetKeys(lhs)),
		Filter(func(key K) bool { return rhs[key] }, GetKeys(rhs)),
	)
}
