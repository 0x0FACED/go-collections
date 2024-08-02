package heaps

import (
	"fmt"
	"testing"

	gocollections "github.com/0x0FACED/go-collections"
	"github.com/stretchr/testify/assert"
)

// compare for max heap
func intComparator_MaxHeap(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

// compare for min heap
func intComparator_MinHeap(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return -1
	} else {
		return 1
	}
}

func TestMaxHeap_Insert(t *testing.T) {
	h := NewHeap(intComparator_MaxHeap)

	h.Insert(3)
	h.Insert(10)
	h.Insert(15)
	h.Insert(22)
	h.Insert(2)
	h.Insert(1)
	h.Insert(56)
	h.Insert(23)
	h.Insert(18)

	fmt.Println(h.elements)

	assert.Equal(t, 9, len(h.elements))
	// root
	assert.Equal(t, 56, h.elements[0])
	// right child
	assert.Equal(t, 22, h.elements[2])
	// left child
	assert.Equal(t, 23, h.elements[1])
}

func TestMinHeap_Insert(t *testing.T) {
	h := NewHeap(intComparator_MinHeap)

	h.Insert(3)
	h.Insert(10)
	h.Insert(15)
	h.Insert(22)
	h.Insert(2)
	h.Insert(1)
	h.Insert(56)
	h.Insert(23)
	h.Insert(18)

	fmt.Println(h.elements)

	assert.Equal(t, 9, len(h.elements))
	// root
	assert.Equal(t, 1, h.elements[0])
	// right child
	assert.Equal(t, 2, h.elements[2])
	// left child
	assert.Equal(t, 3, h.elements[1])
}

func TestMaxHeap_Extract(t *testing.T) {
	h := NewHeap(intComparator_MaxHeap)

	h.Insert(3)
	h.Insert(10)
	h.Insert(5)

	max, err := h.Extract()
	assert.NoError(t, err)
	assert.Equal(t, 10, *max)

	max, err = h.Extract()
	assert.NoError(t, err)
	assert.Equal(t, 5, *max)

	max, err = h.Extract()
	assert.NoError(t, err)
	assert.Equal(t, 3, *max)

	_, err = h.Extract()
	assert.Error(t, err)
	assert.Equal(t, gocollections.ErrEmpty, err.Error())
}

func TestMaxHeap_HeapifyUp(t *testing.T) {
	h := NewHeap(intComparator_MaxHeap)

	h.Insert(3)
	h.Insert(10)
	h.Insert(5)

	assert.Equal(t, 3, h.Size())
	assert.Equal(t, 10, h.elements[0])
	assert.Equal(t, 5, h.elements[2])
	assert.Equal(t, 3, h.elements[1])
}

func TestMaxHeap_HeapifyDown(t *testing.T) {
	h := NewHeap(intComparator_MaxHeap)

	h.Insert(10)
	h.Insert(5)
	h.Insert(3)
	h.Insert(2)

	max, err := h.Extract()
	assert.NoError(t, err)
	assert.Equal(t, 10, *max)

	assert.Equal(t, 5, h.elements[0])
	assert.Equal(t, 3, h.elements[2])
	assert.Equal(t, 2, h.elements[1])
}
