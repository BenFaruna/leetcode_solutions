package stack

import (
	"cmp"
	"errors"
)

var ErrStackEmpty = errors.New("stack is empty")

type Node[T comparable] struct {
	Prev  *Node[T]
	Next  *Node[T]
	Value T
}

type Stack[T cmp.Ordered] struct {
	Head *Node[T]
	Len  int
}

func (s *Stack[T]) Top() (T, error) {
	if s.Len == 0 {
		var zero T
		return zero, ErrStackEmpty
	}
	return s.Head.Value, nil
}

func (s *Stack[T]) Push(x T) {
	currHead := s.Head

	node := &Node[T]{Next: currHead, Value: x}
	s.Head = node
	s.Len++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Len == 0 {
		var zero T
		return zero, ErrStackEmpty
	}

	currHead := s.Head
	s.Head = currHead.Next
	s.Len--

	return currHead.Value, nil
}

func (s *Stack[T]) Min() (T, error) {
	if s.Len == 0 {
		var zero T
		return zero, ErrStackEmpty
	}

	return s.Head.Value, nil
}

func (s *Stack[T]) IterMin() (T, error) {
	if s.Len == 0 {
		var zero T
		return zero, ErrStackEmpty
	}

	minimum := s.Head.Value

	for point := s.Head.Next; point.Next != nil; point = point.Next {
		if point.Value < minimum {
			minimum = point.Value
		}
	}
	return minimum, nil
}
