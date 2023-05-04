package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yotzapon/data-stack/stack"
)

func TestImplStack(t *testing.T) {
	t.Run("happy - push", func(t *testing.T) {
		n := stack.NewStackStructure()

		n.Push(1)
		n.Push("test")
		n.Push(1.5)

		assert.Equal(t, 3, n.CountStackElement())
	})

	t.Run("happy - peek", func(t *testing.T) {
		n := stack.NewStackStructure()

		n.Push(1)
		n.Push("test")
		n.Push(1.5)

		assert.Equal(t, 1.5, n.Peek())
	})

	t.Run("happy - pop", func(t *testing.T) {
		n := stack.NewStackStructure()

		n.Push(1)
		n.Push("test")
		n.Push(1.5)
		n.Pop()

		assert.Equal(t, 2, n.CountStackElement())
	})

	t.Run("isEmpty-false", func(t *testing.T) {
		n := stack.NewStackStructure()

		n.Push(1)
		n.Push("test")
		n.Push(1.5)

		assert.False(t, n.IsEmpty())
	})

	t.Run("isEmpty-true", func(t *testing.T) {
		n := stack.NewStackStructure()

		assert.True(t, n.IsEmpty())
	})
}

func BenchmarkPush(b *testing.B) {
	stack := stack.NewStackStructure()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	stack := stack.NewStackStructure()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkPeek(b *testing.B) {
	stack := stack.NewStackStructure()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}

func BenchmarkEmpty(b *testing.B) {
	stack := stack.NewStackStructure()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.IsEmpty()
	}
}
