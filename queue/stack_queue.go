package queue

import "github.com/0x0FACED/go-collections/stack"

type stackQueue[T comparable] struct {
	// st1 is the main stack that stores all the elements
	st1 stack.Stack[T]

	// st2 is the helper stack used for Enqueue and Dequeue operations
	st2 stack.Stack[T]
}

func NewStackQueue[T comparable]() *stackQueue[T] {
	return &stackQueue[T]{
		st1: stack.NewListStack[T](),
		st2: stack.NewListStack[T](),
	}
}

func (q *stackQueue[T]) Enqueue(item T) error {
	q.st1.Push(item)
	return nil
}

func (q *stackQueue[T]) Dequeue() (*T, error) {
	for !q.st1.IsEmpty() {
		val, err := q.st1.Pop()
		if err != nil {
			return nil, err
		}
		q.st2.Push(*val)
	}

	res, err := q.st2.Pop()

	if err != nil {
		return nil, err
	}

	for !q.st2.IsEmpty() {
		val, err := q.st2.Pop()
		if err != nil {
			return nil, err
		}
		q.st1.Push(*val)
	}
	return res, nil
}

func (q *stackQueue[T]) Peek() (*T, error) {
	for !q.st1.IsEmpty() {
		val, err := q.st1.Pop()
		if err != nil {
			return nil, err
		}
		q.st2.Push(*val)
	}

	res, err := q.st2.Peek()

	if err != nil {
		return nil, err
	}

	for !q.st2.IsEmpty() {
		val, err := q.st2.Pop()
		if err != nil {
			return nil, err
		}
		q.st1.Push(*val)
	}
	return res, nil
}

func (q *stackQueue[T]) IsEmpty() bool {
	return q.st1.IsEmpty() && q.st2.IsEmpty()
}

func (q *stackQueue[T]) Size() int {
	return q.st1.Size()
}

func (q *stackQueue[T]) IsFull() bool {
	return false
}
