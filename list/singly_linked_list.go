package list

import (
	"fmt"
	"reflect"
	"sync"

	gocollections "github.com/0x0FACED/go-collections"
)

type singlyLinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]

	size int

	mu sync.RWMutex
}

func NewSinglyLinked[T comparable]() *singlyLinkedList[T] {
	return &singlyLinkedList[T]{}
}

func (l *singlyLinkedList[T]) Head() *node[T] {
	return l.head
}

func (l *singlyLinkedList[T]) Tail() *node[T] {
	return l.tail
}

func (l *singlyLinkedList[T]) Add(item T) error {
	if l.mu.TryLock() {
		defer l.mu.Unlock()
	}

	node := &node[T]{val: item}
	if l.size == 0 {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = l.tail.next
	}
	l.size++
	return nil
}

func (l *singlyLinkedList[T]) Insert(item T, pos int) error {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.size == 0 || pos == l.size {
		return l.Add(item)
	}

	if pos < 0 || pos > l.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	if pos == 0 {
		node := &node[T]{val: item, next: l.head}
		l.head = node
		l.size++
		return nil
	}

	dummy := l.head
	cnt := 0
	for cnt != pos-1 {
		dummy = dummy.next
		cnt++
	}
	node := &node[T]{val: item, next: dummy.next}
	dummy.next = node
	l.size++
	return nil
}

func (l *singlyLinkedList[T]) RemoveLast() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}
	if l.size == 1 {
		l.head = nil
		l.size--
		return nil
	}
	dummy := l.head
	cnt := 0
	for cnt != l.size-2 {
		dummy = dummy.next
		cnt++
	}
	// 1 3 5 7 8 9 <- before removing
	// 1 3 5 7 8 <- must be after
	// traverse to size-2 pos -> 1 3 5 7 [8] 9
	l.tail = dummy
	dummy.next = nil
	l.size--
	return nil
}

func (l *singlyLinkedList[T]) RemoveVal(item T) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}

	if reflect.DeepEqual(l.head.val, item) {
		l.head = l.head.next
		l.size--
		if l.size == 0 {
			l.tail = nil
		}
		return 0, nil
	}

	dummy := l.head
	cnt := 0
	for dummy.next != nil {
		if reflect.DeepEqual(dummy.next.val, item) {
			dummy.next = dummy.next.next
			if dummy.next == nil {
				l.tail = dummy
			}
			l.size--
			return cnt + 1, nil
		}
		dummy = dummy.next
		cnt++
	}
	return -1, fmt.Errorf(gocollections.ErrNotFound)
}

func (l *singlyLinkedList[T]) RemoveAt(pos int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos > l.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	if pos == 0 {
		l.head = l.head.next
		l.size--
		return nil
	}

	dummy := l.head
	cnt := 0
	for cnt < pos-1 {
		dummy = dummy.next
		cnt++
	}
	dummy.next = dummy.next.next
	l.size--
	return nil
}

func (l *singlyLinkedList[T]) Set(item T, pos int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos > l.size {
		return fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	dummy := l.head
	cnt := 0
	for cnt < pos {
		dummy = dummy.next
		cnt++
	}
	dummy.val = item
	return nil
}

func (l *singlyLinkedList[T]) Get(pos int) (*T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	if pos < 0 || pos > l.size {
		return nil, fmt.Errorf(gocollections.ErrOutOfBounds)
	}

	dummy := l.head
	cnt := 0
	for cnt != pos {
		dummy = dummy.next
		cnt++
	}
	return &dummy.val, nil
}

func (l *singlyLinkedList[T]) GetLast() (*T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return nil, fmt.Errorf(gocollections.ErrEmpty)
	}

	return &l.tail.val, nil
}

func (l *singlyLinkedList[T]) GetPosition(item T) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return -1, fmt.Errorf(gocollections.ErrEmpty)
	}
	dummy := l.head
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

func (l *singlyLinkedList[T]) Size() int {
	return l.size
}

func (l *singlyLinkedList[T]) Clear() error {
	l.head = nil
	l.tail = nil
	l.size = 0
	return nil
}

func (l *singlyLinkedList[T]) Print() {
	if l.mu.TryLock() {
		defer l.mu.Unlock()
	}

	fmt.Println("size:", l.size)
	fmt.Println("data:")
	dummy := l.head
	cnt := 0
	fmt.Print("[ ")
	for dummy != nil {
		fmt.Printf("%v:", cnt)
		fmt.Printf("val=%+v; ", dummy.val)
		dummy = dummy.next
		cnt++
	}
	fmt.Print(" ]\n")
}

func (l *singlyLinkedList[T]) Contains(item T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		return false
	}
	dummy := l.head
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
