package stack

import (
	"errors"
	"testing"
)

func TestMinStack(t *testing.T) {
	t.Run("min stack top test", func(t *testing.T) {
		stack := setupMinStack()

		top, _ := stack.Top()

		if top != 90 {
			t.Logf("top value is %d, want 90", top)
		}
	})

	t.Run("min stack len test", func(t *testing.T) {
		stack := &Stack[int]{}

		if stack.Len() != 0 {
			t.Logf("stack len is %d, want 0", stack.Len())
		}

		stack.Push(5)
		stack.Push(3)
		stack.Push(5)
		stack.Push(5)

		if stack.Len() != 4 {
			t.Logf("stack len is %d, want 1", stack.Len())
		}
	})

	t.Run("min stack push test", func(t *testing.T) {
		stack := &Stack[int]{}

		stack.Push(5)
		stack.Push(3)
		stack.Push(13)

		if stack.Len() != 3 {
			t.Logf("stack len is %d, want 3", stack.Len())
		}

		top, _ := stack.Top()
		if top != 13 {
			t.Logf("stack top is %d, want 13", top)
		}
	})

	t.Run("min stack pop test", func(t *testing.T) {
		stack := setupMinStack()

		for i := 0; i < 5; i++ {
			length := stack.Len()
			top, _ := stack.Top()
			pop, _ := stack.Pop()
			if top != pop {
				t.Logf("want %d, got %d", top, pop)
			}
			if (length - 1) != stack.Len() {
				t.Errorf("old stack length is %d, and new stack length is %d", length, stack.Len())
			}
		}
	})

	t.Run("min stack test top on empty stack", func(t *testing.T) {
		stack := &MinStack[int]{}

		value, err := stack.Top()
		if err == nil {
			t.Errorf("want error, got nil")
		}
		if !errors.Is(err, ErrStackEmpty) {
			t.Errorf("want ErrStackEmpty, got %v", err)
		}
		if value != 0 {
			t.Errorf("want 0, got %d", value)
		}
	})

	t.Run("min stack test pop on empty stack", func(t *testing.T) {
		stack := &Stack[int]{}

		value, err := stack.Pop()
		if err == nil {
			t.Errorf("want error, got nil")
		}
		if !errors.Is(err, ErrStackEmpty) {
			t.Errorf("want ErrStackEmpty, got %v", err)
		}
		if value != 0 {
			t.Errorf("want 0, got %d", value)
		}
	})

	t.Run("min stack min test", func(t *testing.T) {
		stack := setupMinStack()
		minimum, _ := stack.Min()
		if minimum != 10 {
			t.Errorf("want 10, got %d", minimum)
		}
	})

	t.Run("min stack test min on empty stack", func(t *testing.T) {
		stack := &Stack[int]{}

		value, err := stack.Min()
		if err == nil {
			t.Errorf("want error, got nil")
		}
		if !errors.Is(err, ErrStackEmpty) {
			t.Errorf("want ErrStackEmpty, got %v", err)
		}
		if value != 0 {
			t.Errorf("want 0, got %d", value)
		}
	})
}

func setupMinStack() MinStack[int] {
	stack := MinStack[int]{
		Stack[int]{},
		Stack[int]{},
		0,
	}

	for i := 1; i < 10; i++ {
		stack.Push(10 * i)
	}
	return stack
}

func BenchmarkMinStackLen(b *testing.B) {
	stack := setupMinStack()
	for i := 0; i < b.N; i++ {
		stack.Len()
	}
}

func BenchmarkMinStackMin(b *testing.B) {
	stack := setupMinStack()
	for i := 0; i < b.N; i++ {
		stack.Min()
	}
}
