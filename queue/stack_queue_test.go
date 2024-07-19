package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStackQueue_Enqueue(t *testing.T) {
	q := NewStackQueue[int]()

	err := q.Enqueue(1)
	assert.Nil(t, err, "Enqueue should not return an error")

	err = q.Enqueue(2)
	assert.Nil(t, err, "Enqueue should not return an error")

	err = q.Enqueue(3)
	assert.Nil(t, err, "Enqueue should not return an error")

	assert.Equal(t, 3, q.Size(), "Queue size should be 3 after enqueuing 3 elements")
}

func TestStackQueue_Dequeue(t *testing.T) {
	q := NewStackQueue[int]()

	_ = q.Enqueue(1)
	_ = q.Enqueue(2)
	_ = q.Enqueue(3)

	val, err := q.Dequeue()
	assert.Nil(t, err, "Dequeue should not return an error")
	assert.Equal(t, 1, *val, "Dequeued value should be 1")

	val, err = q.Dequeue()
	assert.Nil(t, err, "Dequeue should not return an error")
	assert.Equal(t, 2, *val, "Dequeued value should be 2")

	val, err = q.Dequeue()
	assert.Nil(t, err, "Dequeue should not return an error")
	assert.Equal(t, 3, *val, "Dequeued value should be 3")

	assert.True(t, q.IsEmpty(), "Queue should be empty after dequeuing all elements")
}

func TestStackQueue_EnqueueDequeue(t *testing.T) {
	n := 1000
	q := NewStackQueue[int]()
	for i := 0; i < n; i++ {
		require.NoError(t, q.Enqueue(i))
	}
	assert.Equal(t, false, q.IsFull(), "Stack Queue cant be full (always false)")

	j := 0
	for !q.IsEmpty() {
		val, err := q.Dequeue()
		require.NoError(t, err)
		require.Equal(t, j, *val)
		j++
	}

	for i := 0; i < n; i++ {
		require.NoError(t, q.Enqueue(i))
	}
	assert.Equal(t, false, q.IsFull(), "Stack Queue cant be full (always false)")

	j = 0
	for !q.IsEmpty() {
		val, err := q.Dequeue()
		require.Equal(t, j, *val)
		require.NoError(t, err)
		j++
	}
}

func TestStackQueue_Peek(t *testing.T) {
	q := NewStackQueue[int]()

	_ = q.Enqueue(1)
	_ = q.Enqueue(2)
	_ = q.Enqueue(3)

	val, err := q.Peek()
	assert.Nil(t, err, "Peek should not return an error")
	assert.Equal(t, 1, *val, "Peeked value should be 1")

	_, _ = q.Dequeue()

	val, err = q.Peek()
	assert.Nil(t, err, "Peek should not return an error")
	assert.Equal(t, 2, *val, "Peeked value should be 2")

	_, _ = q.Dequeue()

	val, err = q.Peek()
	assert.Nil(t, err, "Peek should not return an error")
	assert.Equal(t, 3, *val, "Peeked value should be 3")

	_, _ = q.Dequeue()

	val, err = q.Peek()
	assert.NotNil(t, err, "Peek should return an error when queue is empty")
	assert.Nil(t, val, "Peeked value should be nil when queue is empty")
}

func TestStackQueue_IsEmpty(t *testing.T) {
	q := NewStackQueue[int]()

	assert.True(t, q.IsEmpty(), "Queue should be empty initially")

	_ = q.Enqueue(1)
	assert.False(t, q.IsEmpty(), "Queue should not be empty after enqueue")

	_, _ = q.Dequeue()
	assert.True(t, q.IsEmpty(), "Queue should be empty after dequeueing all elements")
}

func TestStackQueue_Size(t *testing.T) {
	q := NewStackQueue[int]()

	assert.Equal(t, 0, q.Size(), "Initial size should be 0")

	_ = q.Enqueue(1)
	_ = q.Enqueue(2)
	_ = q.Enqueue(3)

	assert.Equal(t, 3, q.Size(), "Size should be 3 after enqueuing 3 elements")

	_, _ = q.Dequeue()
	assert.Equal(t, 2, q.Size(), "Size should be 2 after dequeueing one element")

	_, _ = q.Dequeue()
	_, _ = q.Dequeue()

	assert.Equal(t, 0, q.Size(), "Size should be 0 after dequeueing all elements")
}
