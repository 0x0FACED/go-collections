package list

import (
	"fmt"
	"reflect"

	gocollections "github.com/0x0FACED/go-collections"
)

type Comparator[T any] func(a, b T) bool

type arrayList[T comparable] struct {
	items []T

	size        int
	cap         int
	scaleFactor float64
}

func NewArrayList[T comparable]() *arrayList[T] {
	return &arrayList[T]{
		cap:         10,
		items:       make([]T, 10),
		scaleFactor: 0.8,
		size:        0,
	}
}

func (a *arrayList[T]) Add(item T) error {
	if a.size >= int(float64(a.cap)*a.scaleFactor) {
		a.resizeArray()
	}
	a.items[a.size] = item
	a.size++

	return nil
}

func (a *arrayList[T]) Insert(item T, pos int) error {
	if pos < 0 || pos >= a.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}
	if a.size >= int(float64(a.cap)*a.scaleFactor) {
		a.resizeArray()
	}
	newItems := make([]T, a.cap)
	copy(newItems, a.items[:pos])
	copy(newItems[pos+1:], a.items[pos:a.size])
	newItems[pos] = item
	a.items = newItems
	a.size++
	return nil
}

func (a *arrayList[T]) RemoveLast() error {
	if a.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}
	a.items = a.items[:a.size]
	a.size--
	return nil
}

func (a *arrayList[T]) RemoveVal(item T) (int, error) {
	if a.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}

	pos, err := a.findFirst(item)
	if err != nil {
		return -1, err
	}

	err = a.RemoveAt(pos)
	if err != nil {
		return -1, err
	}

	return pos, nil
}

func (a *arrayList[T]) RemoveAt(pos int) error {
	if a.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos >= a.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	copy(a.items[pos:], a.items[pos+1:a.size])
	a.items[a.size-1] = *new(T)
	a.size--
	return nil
}

func (a *arrayList[T]) Set(item T, pos int) error {
	if a.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos >= a.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	a.items[pos] = item
	return nil
}

func (a *arrayList[T]) Get(pos int) (*T, error) {
	if a.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos >= a.size {
		return nil, fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	return &a.items[pos], nil
}

func (a *arrayList[T]) GetLast() (*T, error) {
	if a.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	return &a.items[a.size-1], nil
}

func (a *arrayList[T]) GetPosition(item T) (int, error) {
	if a.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}
	pos, err := a.findFirst(item)
	if err != nil {
		// if err != nil -> err.Error() = "not found"
		return -1, err
	}

	return pos, nil
}

func (a *arrayList[T]) Size() int {
	return a.size
}

func (a *arrayList[T]) Clear() error {
	if a.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}
	a.items = a.items[:0]
	a.size = 0
	a.cap = 10
	return nil
}

func (a *arrayList[T]) Contains(item T) bool {
	if pos, err := a.findFirst(item); pos != -1 && err == nil {
		return true
	}

	return false
}

func (a *arrayList[T]) Sort(compare Comparator[T], sortType int) error {
	if a.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}
	switch sortType {
	case MergeSort:
		a.mergeSort(compare)
	case QuickSort:
		a.quickSort(compare)
	case TimSort:
		a.timSort(compare)
	case BubbleSort:
		a.bubbleSort(compare)
	default:
		return fmt.Errorf("chooose sort type: \n0: timSort \n1: quickSort \n2: mergeSort \n3: bubbleSort")
	}

	return nil
}

func (a *arrayList[T]) resizeArray() {
	newCap := int(float64(a.cap) * 1.5)
	newItems := make([]T, newCap)
	copy(newItems, a.items)
	a.items = newItems
	a.cap = newCap
}

func (a *arrayList[T]) findFirst(item T) (int, error) {
	for i, el := range a.items {
		if reflect.DeepEqual(el, item) {
			return i, nil
		}
	}
	return -1, fmt.Errorf(gocollections.ErrNotFound)
}
