package queue

type node[T comparable] struct {
	val T

	next *node[T]
}

// TODO:
// must be Dynamic Queue interface
//
// and Queue interface
//
// #Queue interface have IsFull() method
//
// unlike Dynamic Queue
type Queue[T comparable] interface {
	Enqueue(item T) error

	Dequeue() (*T, error)

	Peek() (*T, error)

	IsEmpty() bool

	Size() int

	IsFull() bool
}
