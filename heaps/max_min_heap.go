package heaps

import (
	"fmt"
	"sync"

	gocollections "github.com/0x0FACED/go-collections"
)

// this struct is a max OR min heap.
// To make MAX heap -> use compare like this:
//
//	func compare(a, b int) int {
//		if a == b {
//			return 0
//		} else if a < b {
//			return -1
//		} else {
//			return 1
//		}
//	}
//
// To make MIN heap -> use compare like this:
//
//	func compare(a, b int) int {
//		if a == b {
//			return 0
//		} else if a > b {
//			return -1
//		} else {
//			return 1
//		}
//	}
type maxMinHeap[T comparable] struct {
	elements []T

	mu      sync.RWMutex
	compare Comparator[T]
}

func NewHeap[T comparable](compare Comparator[T]) *maxMinHeap[T] {
	return &maxMinHeap[T]{compare: compare, elements: make([]T, 0)}
}

// Insert adds element to heap
func (h *maxMinHeap[T]) Insert(item T) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.elements = append(h.elements, item)
	h.heapifyUp(len(h.elements) - 1)
}

// Extract returns MAX item in Max-Heap
// and MIN in Min-Heap
func (h *maxMinHeap[T]) Extract() (*T, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.extractMax()
}

func (h *maxMinHeap[T]) Peek() (*T, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.elements) == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	return &h.elements[0], nil
}

func (h *maxMinHeap[T]) Size() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.elements)
}

func (h *maxMinHeap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

func (h *maxMinHeap[T]) extractMax() (*T, error) {
	if len(h.elements) == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	max := h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.heapifyDown(0)
	return &max, nil
}

// while LEFT < RIGHT.
// compare method must have this impl for example:
//
//	func compare(a, b int) int {
//		if a == b {
//			return 0
//		} else if a < b {
//			return -1
//		} else {
//			return 1
//		}
//	}
func (h *maxMinHeap[T]) heapifyUp(index int) {
	// we swap new element with hi parent, if newElem > parent
	for h.compare(h.elements[parent(index)], h.elements[index]) < 0 {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *maxMinHeap[T]) heapifyDown(index int) {
	lastIndex := len(h.elements) - 1
	l, r := left(index), right(index)
	var child int
	for l <= lastIndex {
		if l == lastIndex || h.compare(h.elements[l], h.elements[r]) > 0 {
			child = l
		} else {
			child = r
		}
		if h.compare(h.elements[index], h.elements[child]) < 0 {
			h.swap(index, child)
			index = child
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *maxMinHeap[T]) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}
