package queue

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

// sliceQueue is a queue that have capacity.
// So thats slice queue -> circular queue
//
//	rear = (rear + 1) % capacity // in inserting (Enqueue)
//	front = (front + 1) % capacity // in deleting (Dequeue)
//
// if you in NewSliceQueue() didnt push cap as arg
// so queue would have cap = 10
//
// This queue is NOT dynamic.
// If you want Dynamic Slice Queue -> check dymanic_slice_queue.go
type sliceQueue[T comparable] struct {
	queue []T

	front    int // always = 0
	rear     int // tail index of queue
	size     int // actual size of queue
	capacity int // standard capacity = 10
}

// NewSlice Queue creates queue with capacity 10.
//
// Cant change capacity, its const:
//
//		return &sliceQueue[T]{
//		queue:    make([]T, 10),
//		capacity: 10,
//		size:     0,
//		front:    0,
//		rear:     0,
//	}
func NewSliceQueue[T comparable]() *sliceQueue[T] {
	// Here 10 is capacity of queue. You cant change it after create
	return &sliceQueue[T]{
		queue:    make([]T, 10),
		capacity: 10,
		size:     0,
		front:    0,
		rear:     0,
	}
}

func NewSliceQueueWithCap[T comparable](cap int) *sliceQueue[T] {
	return &sliceQueue[T]{
		queue:    make([]T, cap),
		capacity: cap,
		size:     0,
		front:    0,
		rear:     0,
	}
}

func (q *sliceQueue[T]) Enqueue(item T) error {
	if q.IsFull() {
		return fmt.Errorf(gocollections.ErrFull)
	}

	q.queue[q.rear] = item
	q.rear = (q.rear + 1) % q.capacity
	q.size++
	return nil
}

func (q *sliceQueue[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	val := q.queue[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return &val, nil
}

func (q *sliceQueue[T]) Peek() (*T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	val := q.queue[q.front]
	return &val, nil
}

func (q *sliceQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *sliceQueue[T]) Size() int {
	return q.size
}

func (q *sliceQueue[T]) IsFull() bool {
	return q.size == q.capacity
}
