package list

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDoublyLinkedList_Add(t *testing.T) {
	list := NewDoublyLinked[int]()
	err := list.Add(1)
	require.NoError(t, err, "Add should not return an error")

	assert.Equal(t, 1, list.size, "Size should be 1")
	assert.Equal(t, 1, list.head.val, "Head value should be 1")
	assert.Equal(t, 1, list.tail.val, "Tail value should be 1")
	assert.Nil(t, list.head.prev, "Head prev should be nil")
	assert.Nil(t, list.tail.next, "Tail next should be nil")
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	list := NewDoublyLinked[int]()
	_ = list.Add(1)
	_ = list.Add(2)
	err := list.Insert(3, 1)
	require.NoError(t, err, "Insert should not return an error")

	assert.Equal(t, 3, list.size, "Size should be 3")
	assert.Equal(t, 3, list.head.next.val, "Head next value should be 3")
	assert.Equal(t, 2, list.tail.val, "Tail value should be 2")
	assert.Equal(t, list.head.next.next, list.tail, "Head next next should be tail")
	assert.Equal(t, list.tail.prev, list.head.next, "Tail prev should be head next")
}

func TestDoublyLinkedList_RemoveLast(t *testing.T) {
	list := NewDoublyLinked[int]()
	_ = list.Add(1)
	_ = list.Add(2)
	err := list.RemoveLast()
	require.NoError(t, err, "RemoveLast should not return an error")

	assert.Equal(t, 1, list.size, "Size should be 1")
	assert.Equal(t, 1, list.head.val, "Head value should be 1")
	assert.Equal(t, 1, list.tail.val, "Tail value should be 1")
	assert.Nil(t, list.head.prev, "Head prev should be nil")
	assert.Nil(t, list.tail.next, "Tail next should be nil")
}

func TestDoublyLinkedList_RemoveVal(t *testing.T) {
	list := NewDoublyLinked[int]()
	_ = list.Add(1)
	_ = list.Add(2)
	_ = list.Add(3)
	pos, err := list.RemoveVal(2)
	require.NoError(t, err, "RemoveVal should not return an error")
	assert.Equal(t, 1, pos, "Position of removed item should be 1")

	assert.Equal(t, 2, list.size, "Size should be 2")
	assert.Equal(t, 3, list.head.next.val, "Head next value should be 3")
	assert.Equal(t, 3, list.tail.val, "Tail value should be 3")
	assert.Equal(t, list.head, list.tail.prev, "Tail prev should be head")
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	list := NewDoublyLinked[int]()
	_ = list.Add(1)
	_ = list.Add(2)
	_ = list.Add(3)
	err := list.RemoveAt(1)
	require.NoError(t, err, "RemoveAt should not return an error")

	assert.Equal(t, 2, list.size, "Size should be 2")
	assert.Equal(t, 3, list.head.next.val, "Head next value should be 3")
	assert.Equal(t, 3, list.tail.val, "Tail value should be 3")
	assert.Equal(t, list.head, list.tail.prev, "Tail prev should be head")
}

func TestDoublyLinkedList_Get(t *testing.T) {
	list := NewDoublyLinked[int]()
	_ = list.Add(1)
	_ = list.Add(2)
	_ = list.Add(3)
	val, err := list.Get(1)
	require.NoError(t, err, "Get should not return an error")
	assert.Equal(t, 2, *val, "Value at position 1 should be 2")
}

func TestDoublyLinkedList_GetPosition(t *testing.T) {
	list := NewDoublyLinked[int]()
	_ = list.Add(1)
	_ = list.Add(2)
	_ = list.Add(3)
	pos, err := list.GetPosition(2)
	require.NoError(t, err, "GetPosition should not return an error")
	assert.Equal(t, 1, pos, "Position of value 2 should be 1")
}

func TestDoublyLinkedList_Size(t *testing.T) {
	list := NewDoublyLinked[int]()
	assert.Equal(t, 0, list.Size(), "Initial size should be 0")
	_ = list.Add(1)
	_ = list.Add(2)
	_ = list.Add(3)
	_ = list.Add(4)
	_ = list.Add(5)
	assert.Equal(t, 5, list.Size(), "Size should be 5 after adding 5 elements")

	_ = list.RemoveLast()
	assert.Equal(t, 4, list.Size(), "Size should be 4 after deleting the last one")
	_ = list.RemoveAt(0)
	assert.Equal(t, 3, list.Size(), "Size should be 3 after deleting the first one")
	pos, _ := list.RemoveVal(2)
	assert.Equal(t, 0, pos, "Pos of 2 must be 0")
	assert.Equal(t, 2, list.Size(), "Size should be 2 after deleting")
	val, _ := list.Get(0)
	assert.Equal(t, 3, *val, "First val now is 3")
}

func TestDoublyLinkedList_Clear(t *testing.T) {
	list := NewDoublyLinked[int]()
	_ = list.Add(1)
	_ = list.Add(2)
	err := list.Clear()
	require.NoError(t, err, "Clear should not return an error")
	assert.Equal(t, 0, list.size, "Size should be 0 after clearing the list")
	assert.Nil(t, list.head, "Head should be nil after clearing the list")
	assert.Nil(t, list.tail, "Tail should be nil after clearing the list")
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	list := NewDoublyLinked[int]()
	_ = list.Add(1)
	_ = list.Add(2)
	assert.True(t, list.Contains(2), "List should contain 2")
	assert.False(t, list.Contains(3), "List should not contain 3")
}

func TestDoublyLinkedList_Performance(t *testing.T) {
	list := NewDoublyLinked[int]()
	n := 100000

	// Test adding performance
	start := time.Now()
	for i := 0; i < n; i++ {
		_ = list.Add(i)
	}
	end := time.Now()
	times := end.Sub(start).Milliseconds()
	log.Println("Time spent to add to DoublyLinkedList in millis: ", times)
	assert.Equal(t, n, list.Size(), "Size should be equal to n after adding n elements")

	// Test getting performance
	start = time.Now()
	for i := 0; i < n-10; i++ {
		val, err := list.Get(i)
		require.NoError(t, err, "Get should not return an error")
		assert.Equal(t, i, *val, "Value at position should be equal to i")
	}
	end = time.Now()
	times = end.Sub(start).Milliseconds()
	log.Println("Time spent to Get in DoublyLinkedList in millis: ", times)
}
