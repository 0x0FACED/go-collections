package list

import (
	"fmt"
	"reflect"
)

type arrayList[T any] struct {
	items []T

	size        int
	cap         int
	scaleFactor float64
}

// returns empty array list with node == nil and scale factor 0.8
func NewArrayList[T any]() *arrayList[T] {
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
		return fmt.Errorf("insert in out of bounds")
	}
	if a.size >= int(float64(a.cap)*a.scaleFactor) {
		a.resizeArray()
	}
	// example: [0 1 2 3 4 5 6 7 0 0 0 0 0]
	// pos = 3, item = 15
	// will be: [0 1 2 15 3 4 5 6 7 0 0 0 0]
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
		return fmt.Errorf("list is empty")
	}
	a.items = a.items[:a.size]
	a.size--
	return nil
}

func (a *arrayList[T]) RemoveVal(item T) (int, error) {
	if a.size == 0 {
		return -1, fmt.Errorf("list is empty")
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
		return fmt.Errorf("list is empty")
	}

	if pos < 0 || pos >= a.size {
		return fmt.Errorf("removing from out of bounds")
	}

	a.items = append(a.items[:pos], a.items[pos+1:]...)
	a.size--
	return nil
}

func (a *arrayList[T]) Set(item T, pos int) error {
	if a.size == 0 {
		return fmt.Errorf("list is empty")
	}

	if pos < 0 || pos >= a.size {
		return fmt.Errorf("removing from out of bounds")
	}

	a.items[pos] = item
	return nil
}

func (a *arrayList[T]) Get(pos int) (*T, error) {
	if a.size == 0 {
		return nil, fmt.Errorf("list is empty")
	}

	if pos < 0 || pos >= a.size {
		return nil, fmt.Errorf("removing from out of bounds")
	}

	return &a.items[pos], nil
}

func (a *arrayList[T]) GetPosition(item T) (int, error) {
	if a.size == 0 {
		return -1, fmt.Errorf("list is empty")
	}
	pos, err := a.findFirst(item)
	if err != nil {
		return -1, err
	}

	return pos, nil
}

func (a *arrayList[T]) Size() int {
	return a.size
}

func (a *arrayList[T]) Clear() error {
	if a.size == 0 {
		return fmt.Errorf("list is empty")
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
	return -1, fmt.Errorf("not found")
}
