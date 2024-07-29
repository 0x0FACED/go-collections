package queue

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
	"github.com/0x0FACED/go-collections/list"
)

type pq_item[T comparable] struct {
	priority int
	item     T
}

// lpq - List Priority Queue
type lpq[T comparable] struct {
	list list.MutableList[pq_item[T]]
}

// LPQ - List Priority Queue
func NewLPQ[T comparable]() *lpq[T] {
	return &lpq[T]{
		list: list.NewSinglyLinked[pq_item[T]](),
	}
}

func (lpq *lpq[T]) Enqueue(item T, priority int) error {
	if priority < 0 {
		return fmt.Errorf(gocollections.ErrPriority)
	}
	pqItem := pq_item[T]{priority: priority, item: item}
	pos := 0
	for pos < lpq.Size() {
		currItem, _ := lpq.list.Get(pos)
		if priority < currItem.priority {
			break
		}
		pos++
	}
	return lpq.list.Insert(pqItem, pos)
}

func (lpq *lpq[T]) DequeueMax() (*T, error) {
	if lpq.IsEmpty() {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	val, err := lpq.list.GetLast()
	if err != nil {
		return nil, err
	}
	if err := lpq.list.RemoveLast(); err != nil {
		return nil, err
	}
	return &val.item, nil
}

func (lpq *lpq[T]) DequeueMin() (*T, error) {
	if lpq.IsEmpty() {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	val, err := lpq.list.Get(0)
	if err != nil {
		return nil, err
	}
	if err := lpq.list.RemoveAt(0); err != nil {
		return nil, err
	}
	return &val.item, nil
}

func (lpq *lpq[T]) PeekMax() (*T, error) {
	if lpq.IsEmpty() {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	val, err := lpq.list.GetLast()
	if err != nil {
		return nil, err
	}
	return &val.item, nil
}

func (lpq *lpq[T]) PeekMin() (*T, error) {
	if lpq.IsEmpty() {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	val, err := lpq.list.Get(0)
	if err != nil {
		return nil, err
	}
	return &val.item, nil
}

func (lpq *lpq[T]) IsEmpty() bool {
	return lpq.Size() == 0
}

func (lpq *lpq[T]) Size() int {
	return lpq.list.Size()
}

func (lpq *lpq[T]) IsFull() bool {
	return false
}
