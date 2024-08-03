package list

import (
	"fmt"
	"reflect"
	"sync"

	gocollections "github.com/0x0FACED/go-collections"
)

// cdll - Ciruclar Doubly Linked List
type cdll[T any] struct {
	head *dnode[T]
	tail *dnode[T]

	size int

	mu sync.Mutex
}

// CDLL - Doubly Circular Linked List
func NewCDLL[T any]() *cdll[T] {
	return &cdll[T]{}
}

func (d *cdll[T]) Add(item T) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	newNode := &dnode[T]{val: item}
	if d.size == 0 {
		d.head = newNode
		d.tail = newNode
		d.head.next = d.tail
		d.head.prev = d.tail
		d.tail.next = d.head
		d.tail.prev = d.head
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		newNode.next = d.head
		d.tail = newNode
	}
	d.size++
	return nil
}

func (d *cdll[T]) Insert(item T, pos int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos >= d.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	newNode := &dnode[T]{val: item}
	if pos == 0 {
		newNode.next = d.head
		newNode.prev = d.tail
		d.tail.next = newNode
		d.head = newNode
		d.size++
		return nil
	}

	dummy := d.traverseToPosition(pos - 1)
	newNode.next = dummy.next
	newNode.prev = dummy
	dummy.next.prev = newNode
	dummy.next = newNode
	d.size++
	return nil
}

func (d *cdll[T]) RemoveLast() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	d.tail.prev.next = d.head
	d.head.prev = d.tail.prev
	d.tail = d.tail.prev
	d.size--
	return nil
}

func (d *cdll[T]) RemoveVal(item T) (int, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}

	if reflect.DeepEqual(d.head.val, item) {
		d.head.next.prev = d.tail
		d.tail.next = d.head.next
		d.head = d.head.next
		d.size--
		if d.size == 0 {
			d.head = nil
			d.tail = nil
		}
		return 0, nil
	}

	dummy := d.head
	cnt := 0
	for dummy.next != nil {
		if reflect.DeepEqual(dummy.next.val, item) {
			if reflect.DeepEqual(dummy.next, d.tail) {
				d.tail = dummy
			}
			dummy.next = dummy.next.next
			dummy.next.prev = dummy
			d.size--
			if d.size == 0 {
				d.head = nil
				d.tail = nil
			}
			return cnt + 1, nil
		}
		dummy = dummy.next
		cnt++
		if reflect.DeepEqual(dummy, d.head) {
			break
		}
	}

	return -1, fmt.Errorf(gocollections.ErrNotFound)
}

func (d *cdll[T]) RemoveAt(pos int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos >= d.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	if pos == 0 {
		if d.size == 1 {
			d.head = nil
			d.tail = nil
		} else {
			d.head = d.head.next
			d.head.prev = d.tail
			d.tail.next = d.head
		}
		d.size--
		return nil
	}

	dummy := d.traverseToPosition(pos)
	dummy.prev.next = dummy.next
	dummy.next.prev = dummy.prev
	if pos == d.size-1 {
		d.tail = dummy.prev
	}

	d.size--
	return nil
}

func (d *cdll[T]) Set(item T, pos int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos >= d.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	dummy := d.traverseToPosition(pos)
	dummy.val = item
	return nil
}

func (d *cdll[T]) Get(pos int) (*T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos >= d.size {
		return nil, fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	return &d.traverseToPosition(pos).val, nil

}

func (d *cdll[T]) GetLast() (*T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	return &d.tail.val, nil
}

func (d *cdll[T]) GetPosition(item T) (int, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}

	dummy := d.head
	cnt := 0
	for dummy != nil {
		if reflect.DeepEqual(dummy.val, item) {
			return cnt, nil
		}
		dummy = dummy.next
		cnt++
		if reflect.DeepEqual(dummy, d.head) {
			break
		}
	}

	return -1, fmt.Errorf(gocollections.ErrNotFound)
}

func (d *cdll[T]) Size() int {
	return d.size
}

func (d *cdll[T]) Clear() error {
	d.head = nil
	d.tail = nil
	d.size = 0
	return nil
}

func (d *cdll[T]) Contains(item T) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return false
	}

	dummy := d.head
	for dummy != nil {
		if reflect.DeepEqual(dummy.val, item) {
			return true
		}
		dummy = dummy.next
		if reflect.DeepEqual(dummy, d.head) {
			break
		}
	}

	return false
}

func (d *cdll[T]) traverseToPosition(pos int) *dnode[T] {
	if pos == -1 {
		return d.head.prev
	}
	dummy := d.head
	cnt := 0
	for cnt != pos {
		dummy = dummy.next
		cnt++
	}

	return dummy
}
