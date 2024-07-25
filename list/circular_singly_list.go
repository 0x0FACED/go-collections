package list

import (
	"fmt"
	"reflect"
	"sync"

	gocollections "github.com/0x0FACED/go-collections"
)

// Circular Singly Linked List (CSLL) is a linked list where the last node points back to the first node.
//
// It supports operations such as adding, inserting, removing, and retrieving elements.
type csll[T comparable] struct {
	head *node[T]
	tail *node[T]

	size int

	mu sync.Mutex
}

// NewCircularSingly creates a new, empty circular singly linked list.
//
// Returns: pointer to a new csll.
func NewCircularSingly[T comparable]() *csll[T] {
	return &csll[T]{}
}

// Head returns the head of circular singly linked list.
func (c *csll[T]) Head() *node[T] {
	return c.head
}

// Tail returns the tail of circular singly linked list.
func (c *csll[T]) Tail() *node[T] {
	return c.tail
}

// Add adds a new element to the end of the list.
//
// Params:
//   - item: the element to add.
//
// Returns: an error if the operation fails (in this case, it always returns nil XD).
//
// Time Complexity: O(1)
//
// # Space Complexity: O(1)
func (c *csll[T]) Add(item T) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	newNode := &node[T]{val: item}
	if c.size == 0 {
		c.head = newNode
		c.tail = newNode
		c.tail.next = c.head
	} else {
		c.tail.next = newNode
		c.tail = newNode
		c.tail.next = c.head
	}
	c.size++
	return nil
}

// Insert inserts a new element at the specified position in the list.
//
// Params:
//   - item: the element to insert.
//   - pos: the zero-based index at which to insert the element.
//
// Returns: an error if the position is out of bounds or the list is empty.
//
// Time Complexity:
//  1. Best case (inserting to the head): O(1).
//  2. Worst case (inserting to specified position or tail): O(n).
//
// Space Complexity: O(1)
func (c *csll[T]) Insert(item T, pos int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos > c.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	newNode := &node[T]{val: item}
	if pos == 0 {
		newNode.next = c.head
		c.head = newNode
		c.size++
		return nil
	}

	dummy := c.head
	cnt := 0
	for cnt < pos-1 {
		dummy = dummy.next
		cnt++
	}

	newNode.next = dummy.next
	dummy.next = newNode
	c.size++
	return nil
}

// RemoveLast removes the last element from the list.
//
// Returns: an error if the list is empty.
//
// Time Complexity: O(n)
//
// Space Complexity: O(1)
func (c *csll[T]) RemoveLast() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	dummy := c.head
	cnt := 0
	for cnt != c.size-2 {
		dummy = dummy.next
		cnt++
	}
	c.tail = dummy
	dummy.next = c.head
	c.size--

	if c.size == 0 {
		c.tail = nil
	}
	return nil
}

// RemoveVal removes the first occurrence of the specified element from the list.
//
// Params:
//   - item: the element to remove.
//
// Returns: the position of the removed element and an error if the element is not found or the list is empty.
//
// Time Complexity:
//  1. Best case (inserting to the head): O(1).
//  2. Worst case (inserting to the tail): O(n).
//
// Space Complexity: O(1)
func (c *csll[T]) RemoveVal(item T) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}

	if reflect.DeepEqual(c.head.val, item) {
		c.head = c.head.next
		c.tail.next = c.head
		c.size--
		if c.size == 0 {
			c.tail = nil
		}
		return 0, nil
	}

	dummy := c.head
	cnt := 0
	for dummy.next != nil {
		if reflect.DeepEqual(dummy.next.val, item) {
			dummy.next = dummy.next.next
			if dummy.next == c.head {
				c.tail = dummy
			}
			c.size--
			return cnt + 1, nil
		}
		dummy = dummy.next
		cnt++
	}

	return -1, fmt.Errorf(gocollections.ErrNotFound)
}

// RemoveAt removes the element at the specified position in the list.
//
// Params:
//   - pos: the zero-based index of the element to remove.
//
// Returns: an error if the position is out of bounds or the list is empty.
//
// Time Complexity: O(n)
//
// Space Complexity: O(1)
func (c *csll[T]) RemoveAt(pos int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos > c.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	if pos == 0 {
		c.head = c.head.next
		c.tail.next = c.head
		c.size--
		if c.size == 0 {
			c.tail = nil
		}
		return nil
	}

	dummy := c.head
	cnt := 0
	for cnt != pos-1 {
		dummy = dummy.next
		cnt++
	}
	dummy.next = dummy.next.next
	if dummy.next == nil {
		c.tail = dummy
	}
	c.size--
	return nil
}

// Set sets the value of the element at the specified position in the list.
//
// Params:
//   - item: the new value.
//   - pos: the zero-based index of the element to set.
//
// Returns: an error if the position is out of bounds or the list is empty.
//
// Time Complexity: O(n)
//
// Space Complexity: O(1)
func (c *csll[T]) Set(item T, pos int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos > c.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	dummy := c.head
	cnt := 0
	for cnt != pos-1 {
		dummy = dummy.next
		cnt++
	}

	dummy.val = item
	c.size++
	return nil
}

// Get returns the value of the element at the specified position in the list.
//
// Params:
//   - pos: the zero-based index of the element to get.
//
// Returns: a pointer to the value and an error if the position is out of bounds or the list is empty.
//
// Time Complexity: O(n)
//
// Space Complexity: O(1)
func (c *csll[T]) Get(pos int) (*T, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos > c.size {
		return nil, fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	dummy := c.head
	cnt := 0
	for cnt != pos {
		dummy = dummy.next
		if dummy == c.head {
			return nil, fmt.Errorf(gocollections.ErrNotFound)
		}
		cnt++
	}

	return &dummy.val, nil
}

// GetLast returns the last value of the element of the list.
//
// Returns: a pointer to the value and an error if the position is out of bounds or the list is empty.
//
// Time Complexity: O(1)
//
// Space Complexity: O(1)
func (c *csll[T]) GetLast() (*T, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	return &c.tail.val, nil
}

// GetPosition returns the position of the first occurrence of the specified element in the list.
//
// Params:
//   - item: the element to search for.
//
// Returns: the zero-based index of the element and an error if the element is not found or the list is empty.
//
// Time Complexity: O(n)
//
// Space Complexity: O(1)
func (c *csll[T]) GetPosition(item T) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}

	dummy := c.head
	cnt := 0
	for dummy != nil {
		if reflect.DeepEqual(dummy.val, item) {
			return cnt, nil
		}
		dummy = dummy.next
		cnt++
	}

	return -1, fmt.Errorf(gocollections.ErrNotFound)
}

// Size returns the number of elements in the list.
func (c *csll[T]) Size() int {
	return c.size
}

// Clear removes all elements from the list.
// Always error == nil
func (c *csll[T]) Clear() error {
	c.head = nil
	c.tail = nil
	c.size = 0
	return nil
}

// Contains returns true if the list contains the specified element.
//
// Params:
//   - item: the element to search for.
//
// Returns: true if the element is found, false otherwise.
//
// Time Complexity: O(n)
//
// Space Complexity: O(1)
func (c *csll[T]) Contains(item T) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	dummy := c.head
	for dummy != nil {
		if reflect.DeepEqual(dummy.val, item) {
			return true
		}
		dummy = dummy.next
		if dummy == c.head {
			return false
		}
	}

	return false
}
