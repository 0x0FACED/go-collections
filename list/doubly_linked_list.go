package list

import (
	"fmt"
	"reflect"
	"sync"

	gocollections "github.com/0x0FACED/go-collections"
)

type doublyLinkedList[T any] struct {
	head *dnode[T]
	tail *dnode[T]

	size int

	mu sync.Mutex
}

func NewDoublyLinked[T any]() *doublyLinkedList[T] {
	return &doublyLinkedList[T]{}
}

func (d *doublyLinkedList[T]) Head() *dnode[T] {
	return d.head
}

func (d *doublyLinkedList[T]) Tail() *dnode[T] {
	return d.tail
}

func (d *doublyLinkedList[T]) Add(item T) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	newNode := &dnode[T]{val: item}
	if d.size == 0 {
		d.head = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
	}
	d.tail = newNode
	d.size++
	return nil
}

func (d *doublyLinkedList[T]) Insert(item T, pos int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if pos < 0 || pos > d.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	newNode := &dnode[T]{val: item}
	if pos == 0 {
		newNode.next = d.head
		if d.head != nil {
			d.head.prev = newNode
		}
		d.head = newNode
		if d.size == 0 {
			d.tail = newNode
		}
		d.size++
		return nil
	}

	dummy := d.traverseToPosition(pos - 1)
	newNode.prev = dummy
	newNode.next = dummy.next
	if dummy.next != nil {
		dummy.next.prev = newNode
	}
	dummy.next = newNode
	if pos == d.size {
		d.tail = newNode
	}
	d.size++
	return nil
}

func (d *doublyLinkedList[T]) RemoveLast() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if d.size == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.prev
		d.tail.next = nil
	}
	d.size--
	return nil
}

func (d *doublyLinkedList[T]) RemoveVal(item T) (int, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}

	if reflect.DeepEqual(d.head.val, item) {
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		} else {
			d.tail = nil
		}
		d.size--
		return 0, nil
	}

	dummy := d.head
	cnt := 0
	for dummy != nil {
		if reflect.DeepEqual(dummy.val, item) {
			if dummy.prev != nil {
				dummy.prev.next = dummy.next
			}
			if dummy.next != nil {
				dummy.next.prev = dummy.prev
			}
			if dummy == d.tail {
				d.tail = dummy.prev
			}
			d.size--
			return cnt, nil
		}
		dummy = dummy.next
		cnt++
	}

	return -1, fmt.Errorf(gocollections.ErrNotFound)
}

func (d *doublyLinkedList[T]) RemoveAt(pos int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos >= d.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	if pos == 0 {
		if d.head.next != nil {
			d.head = d.head.next
			d.head.prev = nil
		} else {
			d.head = nil
			d.tail = nil
		}
		d.size--
		return nil
	}
	dummy := d.traverseToPosition(pos - 1)
	dummy.next = dummy.next.next
	if dummy.next != nil {
		dummy.next.prev = dummy
	} else {
		d.tail = dummy
	}

	d.size--
	return nil
}

func (d *doublyLinkedList[T]) Set(item T, pos int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if pos < 0 || pos >= d.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	dummy := d.traverseToPosition(pos)
	dummy.val = item
	return nil
}

func (d *doublyLinkedList[T]) Get(pos int) (*T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if pos < 0 || pos >= d.size {
		return nil, fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	return &d.traverseToPosition(pos).val, nil
}

func (d *doublyLinkedList[T]) GetLast() (*T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}
	return &d.tail.val, nil
}

func (d *doublyLinkedList[T]) GetPosition(item T) (int, error) {
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
	}

	return -1, fmt.Errorf(gocollections.ErrNotFound)
}

func (d *doublyLinkedList[T]) Size() int {
	return d.size
}

func (d *doublyLinkedList[T]) Clear() error {
	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}
	d.head = nil
	d.tail = nil
	d.size = 0
	return nil
}

func (d *doublyLinkedList[T]) Contains(item T) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	dummy := d.head
	for dummy != nil {
		if reflect.DeepEqual(dummy.val, item) {
			return true
		}
		dummy = dummy.next
	}

	return false
}

func (d *doublyLinkedList[T]) traverseToPosition(pos int) *dnode[T] {
	if pos < 0 {
		return d.head
	}
	dummy := d.head
	cnt := 0
	for cnt != pos {
		dummy = dummy.next
		cnt++
	}

	return dummy
}
