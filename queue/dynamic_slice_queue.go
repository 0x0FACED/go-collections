package queue

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

// dsq - Dynamic Slice Queue
type dsq[T any] struct {
	queue []T

	front int
	size  int
}

func NewDynamicSliceQueue[T any]() *dsq[T] {
	return &dsq[T]{
		queue: make([]T, 0),
		front: 0,
		size:  0,
	}
}

func (q *dsq[T]) Enqueue(item T) error {
	if q.IsFull() {
		return fmt.Errorf(gocollections.ErrFull)
	}

	q.queue = append(q.queue, item)
	q.size++
	return nil
}

func (q *dsq[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	val := q.queue[q.front]
	q.front++
	q.size--

	if q.front > len(q.queue)/2 {
		q.queue = q.queue[q.front:]
		q.front = 0
	}

	return &val, nil
}

func (q *dsq[T]) Peek() (*T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	return &q.queue[q.front], nil
}

func (q *dsq[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *dsq[T]) Size() int {
	return q.size
}

// since we have dynamic queue, its always false
func (q *dsq[T]) IsFull() bool {
	return false
}
