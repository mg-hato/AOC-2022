package common

// Create a map based on array `arr` such that each element of the array will produce a key-value mapping
// based on `keyf` and `valf` functions
func CreateKeyValueMap[T any, K comparable, V any](arr []T, keyf func(T) K, valf func(T) V) map[K]V {
	var newMap map[K]V = make(map[K]V)
	for _, elem := range arr {
		newMap[keyf(elem)] = valf(elem)
	}
	return newMap
}

// Associate each element of the array with a value
func AssociateWith[K comparable, V any](arr []K, valf func(K) V) map[K]V {
	return CreateKeyValueMap(arr, Identity[K], valf)
}

// Create a key for each element of the array `arr`
func AssociateBy[V any, K comparable](arr []V, keyf func(V) K) map[K]V {
	return CreateKeyValueMap(arr, keyf, Identity[V])
}

// Group elements of `arr` into a map. Keys and values are defined by extraction functions `keyf` and `valf`, respectively.
//
// Key of an element `x` of `arr` is `keyf(x)`
//
// Value of an element `x` of `arr` is `valf(x)`
func GroupBy[T any, K comparable, V any](arr []T, keyf func(T) K, valf func(T) V) map[K][]V {
	var grouped map[K][]V = make(map[K][]V)
	for _, element := range arr {
		key := keyf(element)
		val := valf(element)
		if _, keyAlreadyExists := grouped[key]; !keyAlreadyExists {
			grouped[key] = []V{val}
		} else {
			grouped[key] = append(grouped[key], val)
		}
	}
	return grouped
}

// Get values of the map in an array
func GetValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	var i int = 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Get keys of the map in an array
func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	var i int = 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func MapContains[K comparable, V any](key K) func(map[K]V) bool {
	return func(m map[K]V) bool {
		_, contains_key := m[key]
		return contains_key
	}
}

// Checks if the two maps have the same key-value mappings
func MapEqual[K, V comparable](lhs, rhs map[K]V) bool {
	return MapEqualWith[K](func(left, right V) bool {
		return left == right
	})(lhs, rhs)
}

func MapEqualWith[K comparable, V any](equality_func func(V, V) bool) func(map[K]V, map[K]V) bool {
	return func(lhs, rhs map[K]V) bool {

		if len(lhs) != len(rhs) {
			return false
		}

		for k, v := range lhs {
			if rv, ok := rhs[k]; !ok || !equality_func(v, rv) {
				return false
			}
		}

		return true
	}
}
