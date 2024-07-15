package list

import (
	"fmt"
	"reflect"

	gocollections "github.com/0x0FACED/go-collections"
)

// Circular Singly Linked List - CSLL
type csll[T comparable] struct {
	head *node[T]
	tail *node[T]

	size int
}

func NewCircularSingly[T comparable]() *csll[T] {
	return &csll[T]{}
}

func (c *csll[T]) Add(item T) error {
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

func (c *csll[T]) Insert(item T, pos int) error {
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

func (c *csll[T]) RemoveLast() error {
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

func (c *csll[T]) RemoveVal(item T) (int, error) {
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

func (c *csll[T]) RemoveAt(pos int) error {
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
	c.size--
	return nil
}

func (c *csll[T]) Set(item T, pos int) error {
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

func (c *csll[T]) Get(pos int) (*T, error) {
	if c.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos > c.size {
		return nil, fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	dummy := c.head
	cnt := 0
	for cnt != pos-1 {
		dummy = dummy.next
		cnt++
	}

	return &dummy.val, nil
}

func (c *csll[T]) GetPosition(item T) (int, error) {
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

func (c *csll[T]) Size() int {
	return c.size
}

func (c *csll[T]) Clear() error {
	c.head.next = nil
	c.head = nil
	c.tail = nil
	c.size = 0
	return nil
}

func (c *csll[T]) Contains(item T) bool {
	dummy := c.head
	cnt := 0
	for dummy != nil {
		if reflect.DeepEqual(dummy.val, item) {
			return true
		}
		dummy = dummy.next
		cnt++
	}

	return false
}
