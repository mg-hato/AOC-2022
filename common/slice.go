package common

// Maps function `f` onto every element of `arr` and returns the new mapped array
func Map[T, R any](f func(T) R, arr []T) []R {
	var mappedArray []R = make([]R, len(arr))
	for i, elem := range arr {
		mappedArray[i] = f(elem)
	}
	return mappedArray
}

// Makes a shallow copy of the array
func ShallowCopy[T any](arr []T) []T {
	return Map(Identity[T], arr)
}

// Maps function `f` onto every element of given matrix and returns the new mapped matrix
func MatrixMap[T, R any](f func(T) R, matrix [][]T) [][]R {
	return Map(func(row []T) []R {
		return Map(f, row)
	}, matrix)
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
func MaximumBy[T any](arr []T, lt func(T, T) bool) T {
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
func MinimumBy[T any](arr []T, lt func(T, T) bool) T {
	f := func(lhs, rhs T) T {
		if lt(rhs, lhs) {
			return rhs
		} else {
			return lhs
		}
	}
	return Foldl(f, arr, arr[0])
}

// Returns true iff array contains the given item
func ArrayContains[T comparable](array []T, item T) bool {
	return Any(func(element T) bool { return element == item }, array)
}

// Checks whether the two arrays have the same elements in the same order
func ArrayEqual[T comparable](lhs []T, rhs []T) bool {
	return ArrayEqualWith(func(t1, t2 T) bool { return t1 == t2 })(lhs, rhs)
}

// Checks whether the two array have the same elements in the same order
//
// The elements are compared for equality with a provided `equality_func`
func ArrayEqualWith[T any](equality_func func(T, T) bool) func([]T, []T) bool {
	return func(lhs, rhs []T) bool {
		var size int
		if size = len(lhs); size != len(rhs) {
			return false
		}

		for i := 0; i < size; i++ {
			if !equality_func(lhs[i], rhs[i]) {
				return false
			}
		}
		return true
	}
}

// Checks if two arrays contains exactly the same elements in any order
func ArrayEqualInAnyOrder[T comparable](lhs, rhs []T) bool {
	return MapEqual(CreateSet(lhs, Identity[T]), CreateSet(rhs, Identity[T]))
}

func ArrayEqualInAnyOrderWith[T any](equality_func func(T, T) bool) func([]T, []T) bool {
	return func(lhs, rhs []T) bool {
		if len(lhs) != len(rhs) {
			return false
		}
		rhs_matched := make([]bool, len(rhs))
		enumerated_rhs := Enumerate(rhs)
		for _, lhs_element := range lhs {
			matched_index := IndexOf(enumerated_rhs, func(element Pair[int, T]) bool {
				return !rhs_matched[element.First] && equality_func(lhs_element, element.Second)
			})

			if matched_index == -1 {
				return false
			}
			rhs_matched[matched_index] = true
		}
		return true
	}
}

func Zip[A, B any](arrayA []A, arrayB []B) []Pair[A, B] {
	if arrayA == nil || arrayB == nil {
		return nil
	}

	size := len(arrayA)
	if bsize := len(arrayB); size > bsize {
		size = bsize
	}

	var zipped []Pair[A, B] = make([]Pair[A, B], size)
	for i := 0; i < size; i++ {
		zipped[i] = Pair[A, B]{arrayA[i], arrayB[i]}
	}
	return zipped
}

// Creates an array of length `times` by repeating element `x`
func Repeat[T any](x T, times int) []T {
	times = Max(0, times)
	array := make([]T, times)
	for i := 0; i < times; i++ {
		array[i] = x
	}
	return array
}

// Returns the number of elements in `arr` that satisfy the given predicate
func Count[T any](arr []T, predicate func(T) bool) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if predicate(arr[i]) {
			count++
		}
	}
	return count
}

// Enumerate the array with numbers starting from 0
func Enumerate[T any](arr []T) []Pair[int, T] {
	return EnumerateWithFirstIndex[T](0)(arr)
}

// Enumerate the array with numbers starting from the specified `firstIndex`
func EnumerateWithFirstIndex[T any](first_index int) func([]T) []Pair[int, T] {
	return func(arr []T) []Pair[int, T] {
		return Zip(
			Range(first_index, first_index+len(arr)),
			arr,
		)
	}
}

// Numbers l, l+1, ... r-1, r
func RangeInclusive(l, r int) []int {
	return Range(l, r+1)
}

// Numbers l, l+1, .... r-1
func Range(l, r int) []int {
	l = Min(l, r)
	var arr []int = make([]int, r-l)
	for i := 0; l < r; i++ {
		arr[i] = l
		l++
	}
	return arr
}

// Returns index of first element satisfying predicate or -1 if no such element exists
func IndexOf[T any](arr []T, predicate func(T) bool) int {
	for i := 0; i < len(arr); i++ {
		if predicate(arr[i]) {
			return i
		}
	}
	return -1
}

// Return a reversed version of the array `arr`
func Reverse[T any](arr []T) []T {
	var reversed []T = make([]T, len(arr))
	var j int = len(arr)

	// INVARIANT: i + j = len(arr)
	for i := 0; i < len(arr); i++ {
		j--
		reversed[j] = arr[i]
	}
	return reversed
}
