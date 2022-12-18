package functional

// Identity function
func Identity[T any](x T) T {
	return x
}

// Maps function `f` onto every element of `arr` and returns the new mapped array
func Map[T, R any](f func(T) R, arr []T) []R {
	var mappedArray []R = make([]R, len(arr))
	for i, elem := range arr {
		mappedArray[i] = f(elem)
	}
	return mappedArray
}

// Fold-right
func Foldr[T, R any](f func(T, R) R, arr []T, b R) R {
	var i int = len(arr)
	var result R = b
	for i > 0 {
		i--
		result = f(arr[i], result)
	}
	return result
}

// Fold-left
func Foldl[T, R any](f func(R, T) R, arr []T, b R) R {
	var result R = b
	for _, elem := range arr {
		result = f(result, elem)
	}
	return result
}

// Flattens an array of arrays
func Flatten[T any](arr [][]T) []T {
	var flattenedArray []T = make([]T, Sum(Map(func(x []T) int { return len(x) }, arr)))
	var i int
	for _, arrElem := range arr {
		for _, elem := range arrElem {
			flattenedArray[i] = elem
			i++
		}
	}
	return flattenedArray
}

// Returns the array that contains all elements of `arr` that satisfy the predicate `p`
func Filter[T any](p func(T) bool, arr []T) []T {
	var filteredArray []T = make([]T, len(arr))
	var i int
	for _, elem := range arr {
		if p(elem) {
			filteredArray[i] = elem
			i++
		}
	}
	return filteredArray[:i]
}

// Every element of `arr` is mapped by `f` into an array. Returns flattened mapped array
func FlatMap[T, R any](f func(T) []R, arr []T) []R {
	return Flatten(Map(f, arr))
}

// Performs function `f` on every element of the array `arr`
func ForEach[T any](f func(T), arr []T) {
	for _, elem := range arr {
		f(elem)
	}
}

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

// Take first `n` elements of the array `arr` or take whole array `arr` if `len(arr) < n`
func Take[T any](n int, arr []T) []T {
	if n < 0 {
		n = 0
	}

	if length := len(arr); length < n {
		n = length
	}
	var taken []T = make([]T, n)
	var i int = 0
	for i < n {
		taken[i] = arr[i]
		i++
	}
	return taken
}

// Drop first `n` elements of the array `arr`
func Drop[T any](n int, arr []T) []T {
	if n < 0 {
		n = 0
	}

	var newSize int = len(arr) - n
	if newSize < 0 {
		newSize = 0
	}
	var afterDropping []T = make([]T, newSize)
	for i := 0; i+n < len(arr); i++ {
		afterDropping[i] = arr[i+n]
	}
	return afterDropping
}

// Returns true if there is at least one element in `arr` satisfying the predicate `p`
func Any[T any](p func(T) bool, arr []T) bool {
	var result bool = false
	var i int
	for i < len(arr) && !result {
		result = result || p(arr[i])
		i++
	}
	return result
}

// Returns true only if all the elemenets in `arr` satisfy the predicate `p`
func All[T any](p func(T) bool, arr []T) bool {
	var result bool = true
	var i int
	for i < len(arr) && result {
		result = result && p(arr[i])
		i++
	}
	return result
}

// Returns maximum element defined by less-than function "lt"
func Maximum[T any](arr []T, lt func(T, T) bool) T {
	f := func(lhs, rhs T) T {
		if lt(lhs, rhs) {
			return rhs
		} else {
			return lhs
		}
	}
	return Foldl(f, arr, arr[0])
}

// Returns minimum element defined by less-than function "lt"
func Minimum[T any](arr []T, lt func(T, T) bool) T {
	f := func(lhs, rhs T) T {
		if lt(rhs, lhs) {
			return rhs
		} else {
			return lhs
		}
	}
	return Foldl(f, arr, arr[0])
}

// Checks whether the two arrays have the same elements in the same order
func ArrayEqual[T comparable](lhs []T, rhs []T) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, l := range lhs {
		if l != rhs[i] {
			return false
		}
	}
	return true
}

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func Sum[T number](arr []T) T {
	return Foldl(
		func(lhs, rhs T) T { return lhs + rhs },
		arr, 0,
	)
}
