package stack

import (
	"github.com/0x0FACED/go-collections/list"
)

type listStack[T comparable] struct {
	list list.List[T]
}

func NewListStack[T comparable]() *listStack[T] {
	return &listStack[T]{
		list: list.NewSinglyLinked[T](),
	}
}

func (ls *listStack[T]) Push(item T) {
	ls.list.Add(item)
}

func (ls *listStack[T]) Pop() (*T, error) {
	val, err := ls.list.GetLast()
	if err != nil {
		return nil, err
	}
	ls.list.RemoveLast()
	return val, nil
}

func (ls *listStack[T]) Peek() (*T, error) {
	val, err := ls.list.GetLast()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (ls *listStack[T]) Size() int {
	return ls.list.Size()
}

func (ls *listStack[T]) IsEmpty() bool {
	return ls.Size() == 0
}
