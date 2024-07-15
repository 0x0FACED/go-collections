package list

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

type doublyLinkedList[T comparable] struct {
	head *dnode[T]
	tail *dnode[T]

	size int
}

func NewDoublyLinked[T comparable]() *doublyLinkedList[T] {
	return &doublyLinkedList[T]{}
}

func (d *doublyLinkedList[T]) Add(item T) error {
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
	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}
	if pos < 0 || pos > d.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	newNode := &dnode[T]{val: item}
	if pos == 0 {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
		if d.size == 0 {
			d.tail = newNode
		}
		d.size++
		return nil
	}

	dummy := d.head
	cnt := 0
	for cnt < pos-1 {
		dummy = dummy.next
		cnt++
	}
	newNode.prev = dummy
	newNode.next = dummy.next
	dummy.next = newNode
	dummy.next.prev = newNode
	d.size++
	return nil
}

func (d *doublyLinkedList[T]) RemoveLast() error {
	if d.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if d.size == 1 {
		d.head = nil
	} else {
		d.tail = d.tail.prev
		d.tail.next = nil
	}
	d.size--
	return nil
}

func (d *doublyLinkedList[T]) RemoveVal(item T) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (d *doublyLinkedList[T]) RemoveAt(pos int) error {
	panic("not implemented") // TODO: Implement
}

func (d *doublyLinkedList[T]) Set(item T, pos int) error {
	panic("not implemented") // TODO: Implement
}

func (d *doublyLinkedList[T]) Get(pos int) (*T, error) {
	panic("not implemented") // TODO: Implement
}

func (d *doublyLinkedList[T]) GetPosition(item T) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (d *doublyLinkedList[T]) Size() int {
	panic("not implemented") // TODO: Implement
}

func (d *doublyLinkedList[T]) Clear() error {
	panic("not implemented") // TODO: Implement
}

func (d *doublyLinkedList[T]) Contains(item T) bool {
	panic("not implemented") // TODO: Implement
}
