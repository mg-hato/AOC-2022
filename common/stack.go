package common

import "errors"

type stack[T any] struct {
	stack_arr []T
}

func Stack[T any](elements ...T) *stack[T] {
	return &stack[T]{ShallowCopy(elements)}
}

// Pushes elements on top of the stack
func (s *stack[T]) Push(element T, elements ...T) {
	s.stack_arr = append(s.stack_arr, element)
	s.stack_arr = append(s.stack_arr, elements...)
}

// Pops an element from the top of the stack if there is one, otherwise returns an error
func (s *stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return GetZero[T](), stack_empty_error()
	} else {
		popped_element := s.stack_arr[s.Size()-1]
		s.stack_arr = s.stack_arr[:s.Size()-1]
		return popped_element, nil
	}
}

func (s *stack[T]) Size() int {
	return len(s.stack_arr)
}

func (s *stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Returns element on the top of the stack
func (s *stack[T]) Top() (T, error) {
	if s.IsEmpty() {
		return GetZero[T](), stack_empty_error()
	} else {
		return s.stack_arr[s.Size()-1], nil
	}
}

// Returns stack elements from top to bottom
func (s *stack[T]) GetAsArray() []T {
	return Reverse(s.stack_arr)
}

// Invisible methods / functions

func stack_empty_error() error {
	return errors.New("stack is empty")
}
