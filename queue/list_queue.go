package queue

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

type listQueue[T comparable] struct {
	head *node[T]
	tail *node[T]

	size     int
	capacity int // default capacity = 10
}

func NewListQueue[T comparable]() *listQueue[T] {
	return &listQueue[T]{
		size:     0,
		capacity: 10,
	}
}

func NewListQueueWithCap[T comparable](cap int) *listQueue[T] {
	return &listQueue[T]{
		size:     0,
		capacity: cap,
	}
}
func (q *listQueue[T]) Enqueue(item T) error {
	if q.IsFull() {
		return fmt.Errorf(gocollections.ErrFull)
	}
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
func (q *listQueue[T]) Dequeue() (*T, error) {
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

func (q *listQueue[T]) Peek() (*T, error) {
	if q.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	return &q.head.val, nil
}

func (q *listQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *listQueue[T]) Size() int {
	return q.size
}

func (q *listQueue[T]) IsFull() bool {
	return q.size == q.capacity
}
