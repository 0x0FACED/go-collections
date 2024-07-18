package queue

import (
	"errors"
	"testing"

	gocollections "github.com/0x0FACED/go-collections"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArrayQueue_Enqueue(t *testing.T) {
	q := NewSliceQueue[int]()
	require.NoError(t, q.Enqueue(1))
	require.NoError(t, q.Enqueue(2))
	require.NoError(t, q.Enqueue(3))

	assert.Equal(t, 3, q.Size(), "Size should be 3")
}

func TestArrayQueue_EnqueueQequeue(t *testing.T) {
	n := 10000000
	q := NewSliceQueueWithCap[int](n)
	for i := 0; i < n; i++ {
		require.NoError(t, q.Enqueue(i))
	}
	assert.Equal(t, true, q.IsFull(), "Q must be full")
	assert.Error(t, errors.New(gocollections.ErrFull), q.Enqueue(-1))

	j := 0
	for !q.IsEmpty() {
		val, err := q.Dequeue()
		require.Equal(t, j, *val)
		require.NoError(t, err)
		j++
	}

	for i := 0; i < n; i++ {
		require.NoError(t, q.Enqueue(i))
	}
	assert.Equal(t, true, q.IsFull(), "Q must be full")
	assert.Error(t, errors.New(gocollections.ErrFull), q.Enqueue(-1))

	j = 0
	for !q.IsEmpty() {
		val, err := q.Dequeue()
		require.Equal(t, j, *val)
		require.NoError(t, err)
		j++
	}
}

func TestArrayQueue_Dequeue(t *testing.T) {
	q := NewSliceQueue[int]()
	_ = q.Enqueue(1)
	_ = q.Enqueue(2)
	_ = q.Enqueue(3)

	val, err := q.Dequeue()
	require.NoError(t, err)
	assert.Equal(t, 1, *val, "Dequeued value should be 1")

	val, err = q.Dequeue()
	require.NoError(t, err)
	assert.Equal(t, 2, *val, "Dequeued value should be 2")

	val, err = q.Dequeue()
	require.NoError(t, err)
	assert.Equal(t, 3, *val, "Dequeued value should be 3")

	_, err = q.Dequeue()
	assert.Error(t, err, "Dequeue from empty queue should return an error")
}

func TestArrayQueue_Peek(t *testing.T) {
	q := NewSliceQueue[int]()
	_ = q.Enqueue(1)
	_ = q.Enqueue(2)

	val, err := q.Peek()
	require.NoError(t, err)
	assert.Equal(t, 1, *val, "Peek value should be 1")

	_, _ = q.Dequeue()
	val, err = q.Peek()
	require.NoError(t, err)
	assert.Equal(t, 2, *val, "Peek value should be 2")
}

func TestArrayQueue_IsEmpty(t *testing.T) {
	q := NewSliceQueue[int]()
	assert.True(t, q.IsEmpty(), "Queue should be empty initially")

	_ = q.Enqueue(1)
	assert.False(t, q.IsEmpty(), "Queue should not be empty after enqueue")
}

func TestArrayQueue_Size(t *testing.T) {
	q := NewSliceQueue[int]()
	assert.Equal(t, 0, q.Size(), "Initial size should be 0")

	_ = q.Enqueue(1)
	_ = q.Enqueue(2)
	assert.Equal(t, 2, q.Size(), "Size should be 2 after adding 2 elements")

	_, _ = q.Dequeue()
	assert.Equal(t, 1, q.Size(), "Size should be 1 after one dequeue")
}
