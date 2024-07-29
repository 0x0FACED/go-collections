package queue

import "github.com/0x0FACED/go-collections/list"

type deque[T comparable] struct {
	// list is doubly linked list
	list list.List[T]
}

func NewDeque[T comparable]() *deque[T] {
	return &deque[T]{
		list: list.NewDoublyLinked[T](),
	}
}

func (d *deque[T]) FrontEnqueue(item T) error {
	return d.list.Insert(item, 0)
}

func (d *deque[T]) FrontDequeue() (*T, error) {
	val, err := d.list.Get(0)
	if err != nil {
		return nil, err
	}
	return val, d.list.RemoveAt(0)
}

func (d *deque[T]) FrontPeek() (*T, error) {
	return d.list.Get(0)
}

func (d *deque[T]) Enqueue(item T) error {
	return d.list.Add(item)
}

func (d *deque[T]) Dequeue() (*T, error) {
	val, err := d.list.GetLast()
	if err != nil {
		return nil, err
	}
	return val, d.list.RemoveLast()
}

func (d *deque[T]) Peek() (*T, error) {
	return d.list.GetLast()
}

func (d *deque[T]) IsEmpty() bool {
	return d.list.Size() == 0
}

func (d *deque[T]) Size() int {
	return d.list.Size()
}

// Deque is dynamic, cant be full
func (d *deque[T]) IsFull() bool {
	return false
}
