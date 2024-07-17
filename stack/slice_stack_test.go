package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceStack_Push(t *testing.T) {
	s := NewSliceStack[int]()
	s.Push(1)
	assert.Equal(t, 1, s.Size(), "Size should be 1 after one push")
	s.Push(2)
	assert.Equal(t, 2, s.Size(), "Size should be 2 after two pushes")
}

func TestSliceStack_Pop(t *testing.T) {
	s := NewSliceStack[int]()
	s.Push(1)
	s.Push(2)
	val, err := s.Pop()
	require.NoError(t, err, "Pop should not return an error")
	assert.Equal(t, 2, *val, "Pop should return the last pushed value")
	assert.Equal(t, 1, s.Size(), "Size should be 1 after one pop")
}

func TestSliceStack_Peek(t *testing.T) {
	s := NewSliceStack[int]()
	s.Push(1)
	s.Push(2)
	val, err := s.Peek()
	require.NoError(t, err, "Peek should not return an error")
	assert.Equal(t, 2, *val, "Peek should return the last pushed value without removing it")
	assert.Equal(t, 2, s.Size(), "Size should remain 2 after peek")
}

func TestSliceStack_IsEmpty(t *testing.T) {
	s := NewSliceStack[int]()
	assert.True(t, s.IsEmpty(), "New stack should be empty")
	s.Push(1)
	assert.False(t, s.IsEmpty(), "Stack should not be empty after push")
	s.Pop()
	assert.True(t, s.IsEmpty(), "Stack should be empty after pop")
}

func TestSliceStack_Size(t *testing.T) {
	s := NewSliceStack[int]()
	assert.Equal(t, 0, s.Size(), "New stack should have size 0")
	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, s.Size(), "Stack size should be 2 after two pushes")
	s.Pop()
	assert.Equal(t, 1, s.Size(), "Stack size should be 1 after one pop")
}
