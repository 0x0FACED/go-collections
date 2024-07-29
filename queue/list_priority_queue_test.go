package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLPQ_Enqueue(t *testing.T) {
	pq := NewLPQ[int]()
	fmt.Println("Started")
	assert.NoError(t, pq.Enqueue(5, 2))
	fmt.Println("1")
	assert.NoError(t, pq.Enqueue(3, 1))
	fmt.Println("2")
	assert.NoError(t, pq.Enqueue(9, 3))
	fmt.Println("3")
	assert.Equal(t, 3, pq.Size())
}

func TestLPQ_DequeueMax(t *testing.T) {
	pq := NewLPQ[int]()
	pq.Enqueue(5, 2)
	pq.Enqueue(3, 1)
	pq.Enqueue(9, 3)

	item, err := pq.DequeueMax()
	assert.NoError(t, err)
	assert.Equal(t, 9, *item)
	assert.Equal(t, 2, pq.Size())

	item, err = pq.DequeueMax()
	assert.NoError(t, err)
	assert.Equal(t, 5, *item)
	assert.Equal(t, 1, pq.Size())

	item, err = pq.DequeueMax()
	assert.NoError(t, err)
	assert.Equal(t, 3, *item)
	assert.Equal(t, 0, pq.Size())

	_, err = pq.DequeueMax()
	assert.Error(t, err)
}

func TestLPQ_DequeueMin(t *testing.T) {
	pq := NewLPQ[int]()
	pq.Enqueue(5, 2)
	pq.Enqueue(3, 1)
	pq.Enqueue(9, 3)

	item, err := pq.DequeueMin()
	assert.NoError(t, err)
	assert.Equal(t, 3, *item)
	assert.Equal(t, 2, pq.Size())

	item, err = pq.DequeueMin()
	assert.NoError(t, err)
	assert.Equal(t, 5, *item)
	assert.Equal(t, 1, pq.Size())

	item, err = pq.DequeueMin()
	assert.NoError(t, err)
	assert.Equal(t, 9, *item)
	assert.Equal(t, 0, pq.Size())

	_, err = pq.DequeueMin()
	assert.Error(t, err)
}

func TestLPQ_PeekMax(t *testing.T) {
	pq := NewLPQ[int]()
	pq.Enqueue(5, 2)
	pq.Enqueue(3, 1)
	pq.Enqueue(9, 3)

	item, err := pq.PeekMax()
	assert.NoError(t, err)
	assert.Equal(t, 9, *item)
	assert.Equal(t, 3, pq.Size())
}

func TestLPQ_PeekMin(t *testing.T) {
	pq := NewLPQ[int]()
	pq.Enqueue(5, 2)
	pq.Enqueue(3, 1)
	pq.Enqueue(9, 3)

	item, err := pq.PeekMin()
	assert.NoError(t, err)
	assert.Equal(t, 3, *item)
	assert.Equal(t, 3, pq.Size())
}

func TestLPQ_IsEmpty(t *testing.T) {
	pq := NewLPQ[int]()
	assert.True(t, pq.IsEmpty())

	pq.Enqueue(5, 2)
	assert.False(t, pq.IsEmpty())

	pq.DequeueMax()
	pq.DequeueMax()
	assert.True(t, pq.IsEmpty())
}

func TestLPQ_Size(t *testing.T) {
	pq := NewLPQ[int]()
	assert.Equal(t, 0, pq.Size())

	pq.Enqueue(5, 2)
	assert.Equal(t, 1, pq.Size())

	pq.Enqueue(3, 1)
	assert.Equal(t, 2, pq.Size())
}
