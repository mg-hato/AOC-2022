package common

import "errors"

type queue[T any] struct {
	enqueue_arr []T
	dequeue_arr []T
}

func Queue[T any](elements ...T) *queue[T] {
	return &queue[T]{enqueue_arr: ShallowCopy(elements)}
}

// Enqueues provided elements in the queue
func (q *queue[T]) Enqueue(element T, elements ...T) {
	q.enqueue_arr = append(q.enqueue_arr, element)
	q.enqueue_arr = append(q.enqueue_arr, elements...)
}

// Dequeues an element from the queue, returns an error if queue is empty
func (q *queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return GetZero[T](), queue_empty_error()
	}

	if len(q.dequeue_arr) == 0 {
		q.move_elements_to_dequeue_array()
	}

	dequeued := q.dequeue_arr[len(q.dequeue_arr)-1]
	q.dequeue_arr = q.dequeue_arr[:len(q.dequeue_arr)-1]
	return dequeued, nil
}

func (q *queue[T]) Size() int {
	return len(q.enqueue_arr) + len(q.dequeue_arr)
}

func (q *queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Returns the next element for dequeueing without changing the queue size
func (q *queue[T]) Next() (T, error) {
	if q.IsEmpty() {
		return GetZero[T](), queue_empty_error()
	}

	if len(q.dequeue_arr) == 0 {
		q.move_elements_to_dequeue_array()
	}

	return q.dequeue_arr[len(q.dequeue_arr)-1], nil
}

// Empties queue
func (q *queue[T]) ClearAll() {
	q.dequeue_arr = nil
	q.enqueue_arr = nil
}

// Get all elements of the queue, they are ordered by their dequeueing order
func (q *queue[T]) GetAsArray() []T {
	q.move_elements_to_dequeue_array()
	return Reverse(q.dequeue_arr)
}

// Invisible methods / functions

func queue_empty_error() error {
	return errors.New("queue is empty")
}

func (q *queue[T]) move_elements_to_dequeue_array() {
	q.dequeue_arr = append(q.dequeue_arr, Reverse(q.enqueue_arr)...)
	q.enqueue_arr = nil
}
