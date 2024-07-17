package stack

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

type sliceStack[T comparable] struct {
	elements []T
}

func NewSliceStack[T comparable]() *sliceStack[T] {
	return &sliceStack[T]{
		elements: []T{},
	}
}

func (ss *sliceStack[T]) Push(item T) {
	ss.elements = append(ss.elements, item)
}

func (ss *sliceStack[T]) Pop() (*T, error) {
	if len(ss.elements) == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	idx := len(ss.elements) - 1
	val := ss.elements[idx]
	ss.elements = ss.elements[:idx]

	return &val, nil
}

func (ss *sliceStack[T]) Peek() (*T, error) {
	if len(ss.elements) == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	return &ss.elements[len(ss.elements)-1], nil
}

func (ss *sliceStack[T]) Size() int {
	return len(ss.elements)
}

func (ss *sliceStack[T]) IsEmpty() bool {
	return len(ss.elements) == 0
}
