package queue

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

// dlq - Dynamic List Queue
type dlq[T comparable] struct {
	head *node[T]
	tail *node[T]

	size int
}

func NewDynamicListQueue[T comparable]() *dlq[T] {
	return &dlq[T]{
		size: 0,
	}
}

func (q *dlq[T]) Enqueue(item T) error {
	newNode := &node[T]{val: item}
	if q.size == 0 {
		q.head = newNode
		q.head.next = q.tail
	} else {
		q.tail.next = newNode
	}
	q.tail = newNode
	q.size++
	return nil
}
func (q *dlq[T]) Dequeue() (*T, error) {
	if q.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	val := q.head.val
	q.head = q.head.next
	q.size--
	if q.size == 0 {
		q.head = nil
		q.tail = nil
	}
	return &val, nil
}

func (q *dlq[T]) Peek() (*T, error) {
	if q.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	return &q.head.val, nil
}

func (q *dlq[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *dlq[T]) Size() int {
	return q.size
}

// always false
func (q *dlq[T]) IsFull() bool {
	return false
}
