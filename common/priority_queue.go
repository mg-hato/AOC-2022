package common

import "errors"

type priority_queue[T any] struct {
	elements []T
	cmp      func(T, T) bool
}

// Makes a priority queue with given comparator function to perform prioritisation.
// Given two elements x, y: x has higher priority iff comparator(x, y) returns true.
// The element on top is always of the highest priority
func PriorityQueue[T any](comparator func(T, T) bool) *priority_queue[T] {
	return &priority_queue[T]{cmp: comparator}
}

// Adds multiple elements
func (pq *priority_queue[T]) MultiAdd(elements ...T) {
	ForEach(pq.Add, elements)
}

func (pq *priority_queue[T]) Add(element T) {
	pq.elements = append(pq.elements, element)

	// Move the current upwards
	current := pq.Size() - 1
	for current != 0 && !pq.in_order(parent(current), current) {
		pq.swap(parent(current), current)
		current = parent(current)
	}
}

func (pq *priority_queue[T]) Pop() (T, error) {
	if pq.IsEmpty() {
		return GetZero[T](), priority_queue_is_empty_error()
	}
	popped := pq.elements[0]

	pq.elements[0] = pq.elements[pq.Size()-1]
	pq.elements = pq.elements[:pq.Size()-1]

	current := 0
	stabilised := false
	for !stabilised {
		// children indices that invalidate the invariant
		indices_for_swapping := Filter(
			func(i int) bool {
				return i < pq.Size() && !pq.in_order(current, i)
			},
			[]int{left(current), right(current)},
		)

		if len(indices_for_swapping) == 0 {
			stabilised = true
		} else {
			next := MinimumBy(indices_for_swapping, pq.in_order)
			pq.swap(next, current)
			current = next
		}
	}
	return popped, nil
}

func (pq *priority_queue[T]) Size() int {
	return len(pq.elements)
}

func (pq *priority_queue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

// Invisible methods

// returns true if element[i] has higher or equal priority than element[j]
func (pq *priority_queue[T]) in_order(i, j int) bool {
	return !pq.cmp(pq.elements[j], pq.elements[i])
}

func (pq *priority_queue[T]) swap(i, j int) {
	pq.elements[i], pq.elements[j] = pq.elements[j], pq.elements[i]
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func priority_queue_is_empty_error() error {
	return errors.New("priority queue is empty")
}
