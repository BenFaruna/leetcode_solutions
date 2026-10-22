package stack

import (
	"cmp"
	"errors"
)

type MinStack[T cmp.Ordered] struct {
	mainStack Stack[T]
	minStack  Stack[T]
	len       int
}

func (m *MinStack[T]) Top() (T, error) {
	return m.mainStack.Top()
}

func (m *MinStack[T]) Pop() (T, error) {
	pop, err := m.mainStack.Pop()
	if err != nil {
		return pop, err
	}

	m.len--
	minTop, _ := m.minStack.Top()
	if minTop == pop {
		_, _ = m.minStack.Pop()
	}
	return pop, nil
}

func (m *MinStack[T]) Push(val T) {
	if minTop, err := m.minStack.Top(); errors.Is(err, ErrStackEmpty) || minTop >= val {
		m.minStack.Push(val)
	}
	m.mainStack.Push(val)
	m.len++
}

func (m *MinStack[T]) Len() int {
	return m.len
}

func (m *MinStack[T]) Min() (T, error) {
	return m.minStack.Top()
}
