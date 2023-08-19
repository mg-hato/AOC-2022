package common

import "errors"

type Stack[T any] struct {
	stack_arr []T
}

func MakeStack[T any](elements ...T) *Stack[T] {
	return &Stack[T]{ShallowCopy(elements)}
}

// Pushes elements on top of the stack
func (s *Stack[T]) Push(element T, elements ...T) {
	s.stack_arr = append(s.stack_arr, element)
	s.stack_arr = append(s.stack_arr, elements...)
}

// Pops an element from the top of the stack if there is one, otherwise returns an error
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return GetZero[T](), stack_empty_error()
	} else {
		popped_element := s.stack_arr[s.Size()-1]
		s.stack_arr = s.stack_arr[:s.Size()-1]
		return popped_element, nil
	}
}

func (s *Stack[T]) Size() int {
	return len(s.stack_arr)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Returns element on the top of the stack
func (s *Stack[T]) Top() (T, error) {
	if s.IsEmpty() {
		return GetZero[T](), stack_empty_error()
	} else {
		return s.stack_arr[s.Size()-1], nil
	}
}

// Returns stack elements from top to bottom
func (s *Stack[T]) GetAsArray() []T {
	return Reverse(s.stack_arr)
}

// Invisible methods / functions

func stack_empty_error() error {
	return errors.New("stack is empty")
}
