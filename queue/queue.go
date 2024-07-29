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

type doubleEnded[T comparable] interface {
	FrontEnqueue(item T) error
	FrontDequeue() (*T, error)
	FrontPeek() (*T, error)
}

type Deque[T comparable] interface {
	Queue[T]
	doubleEnded[T]
}

// pq is the interface for Priority Queue.
//
// Because pq methods are different from the others.
//
// Enqueue uses priority to insert the item and balance struct.
//
// DequeueMax() deletes and returns item with max priority.
//
// DequeueMin() with min priority.
//
// PeekMax() same as DequeueMax(), but doesn't remove.
//
// PeekMin() same as DequeueMin(), but doesn't remove.
type pq[T comparable] interface {
	Enqueue(item T, priority int) error

	DequeueMax() (*T, error)
	DequeueMin() (*T, error)

	PeekMax() (*T, error)
	PeekMin() (*T, error)

	IsEmpty() bool

	Size() int
}
