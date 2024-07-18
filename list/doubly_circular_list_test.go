package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to create a new list and add initial elements
func createCDLLWithElements[T comparable](elements ...T) *cdll[T] {
	list := &cdll[T]{}
	for _, el := range elements {
		list.Add(el)
	}
	return list
}

func TestCDLL_Add(t *testing.T) {
	list := &cdll[int]{}
	err := list.Add(1)
	require.NoError(t, err, "Add should not return an error")
	assert.Equal(t, 1, list.Size(), "Size should be 1 after one addition")
	assert.Equal(t, 1, list.head.val, "Head value should be 1")
	assert.Equal(t, 1, list.tail.val, "Tail value should be 1")
	assert.Equal(t, list.head, list.tail, "Head and tail should be the same for single element list")
	assert.Equal(t, list.head.next, list.head, "Next of head should point to head itself")
	assert.Equal(t, list.head.prev, list.head, "Prev of head should point to head itself")
}

func TestCDLL_RemoveAt_Head(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	err := list.RemoveAt(0)
	require.NoError(t, err, "RemoveAt should not return an error")
	assert.Equal(t, 2, list.Size(), "Size should be 2 after removing head")
	assert.Equal(t, 2, list.head.val, "New head value should be 2")
	assert.Equal(t, 3, list.tail.val, "Tail value should remain 3")
	assert.Equal(t, list.head.prev, list.tail, "Head's prev should be tail")
	assert.Equal(t, list.tail.next, list.head, "Tail's next should be head")
}

func TestCDLL_RemoveAt_Tail(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	err := list.RemoveAt(2)
	require.NoError(t, err, "RemoveAt should not return an error")
	assert.Equal(t, 2, list.Size(), "Size should be 2 after removing tail")
	assert.Equal(t, 1, list.head.val, "Head value should remain 1")
	assert.Equal(t, 2, list.tail.val, "New tail value should be 2")
	assert.Equal(t, list.head.prev, list.tail, "Head's prev should be new tail")
	assert.Equal(t, list.tail.next, list.head, "Tail's next should be head")
}

func TestCDLL_RemoveAt_Middle(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	err := list.RemoveAt(1)
	require.NoError(t, err, "RemoveAt should not return an error")
	assert.Equal(t, 2, list.Size(), "Size should be 2 after removing middle element")
	assert.Equal(t, 1, list.head.val, "Head value should remain 1")
	assert.Equal(t, 3, list.tail.val, "Tail value should remain 3")
	assert.Equal(t, list.head.next, list.tail, "Head's next should be tail")
	assert.Equal(t, list.tail.prev, list.head, "Tail's prev should be head")
}

func TestCDLL_RemoveAt_SingleElement(t *testing.T) {
	list := createCDLLWithElements(1)
	err := list.RemoveAt(0)
	require.NoError(t, err, "RemoveAt should not return an error")
	assert.Equal(t, 0, list.Size(), "Size should be 0 after removing the only element")
	assert.Nil(t, list.head, "Head should be nil after removing the only element")
	assert.Nil(t, list.tail, "Tail should be nil after removing the only element")
}

func TestCDLL_RemoveVal_Head(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	pos, err := list.RemoveVal(1)
	require.NoError(t, err, "RemoveVal should not return an error")
	assert.Equal(t, 0, pos, "Position of removed item should be 0")
	assert.Equal(t, 2, list.Size(), "Size should be 2 after removing head value")
	assert.Equal(t, 2, list.head.val, "New head value should be 2")
	assert.Equal(t, 3, list.tail.val, "Tail value should remain 3")
	assert.Equal(t, list.head.prev, list.tail, "Head's prev should be tail")
	assert.Equal(t, list.tail.next, list.head, "Tail's next should be head")
}

func TestCDLL_RemoveVal_Tail(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	pos, err := list.RemoveVal(3)
	require.NoError(t, err, "RemoveVal should not return an error")
	assert.Equal(t, 2, pos, "Position of removed item should be 2")
	assert.Equal(t, 2, list.Size(), "Size should be 2 after removing tail value")
	assert.Equal(t, 1, list.head.val, "Head value should remain 1")
	assert.Equal(t, 2, list.tail.val, "New tail value should be 2")
	assert.Equal(t, list.head.prev, list.tail, "Head's prev should be new tail")
	assert.Equal(t, list.tail.next, list.head, "Tail's next should be head")
}

func TestCDLL_RemoveVal_Middle(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	pos, err := list.RemoveVal(2)
	require.NoError(t, err, "RemoveVal should not return an error")
	assert.Equal(t, 1, pos, "Position of removed item should be 1")
	assert.Equal(t, 2, list.Size(), "Size should be 2 after removing middle value")
	assert.Equal(t, 1, list.head.val, "Head value should remain 1")
	assert.Equal(t, 3, list.tail.val, "Tail value should remain 3")
	assert.Equal(t, list.head.next, list.tail, "Head's next should be tail")
	assert.Equal(t, list.tail.prev, list.head, "Tail's prev should be head")
}

func TestCDLL_RemoveVal_SingleElement(t *testing.T) {
	list := createCDLLWithElements(1)
	pos, err := list.RemoveVal(1)
	require.NoError(t, err, "RemoveVal should not return an error")
	assert.Equal(t, 0, pos, "Position of removed item should be 0")
	assert.Equal(t, 0, list.Size(), "Size should be 0 after removing the only element")
	assert.Nil(t, list.head, "Head should be nil after removing the only element")
	assert.Nil(t, list.tail, "Tail should be nil after removing the only element")
}

func TestCDLL_Insert_Head(t *testing.T) {
	list := createCDLLWithElements(2, 3)
	err := list.Insert(1, 0)
	require.NoError(t, err, "Insert should not return an error")
	assert.Equal(t, 3, list.Size(), "Size should be 3 after insertion at head")
	assert.Equal(t, 1, list.head.val, "Head value should be 1")
	assert.Equal(t, 2, list.head.next.val, "Next value after head should be 2")
	assert.Equal(t, 3, list.tail.val, "Tail value should remain 3")
}

func TestCDLL_Insert_Tail(t *testing.T) {
	list := createCDLLWithElements(1, 2)
	err := list.Insert(3, 1)
	require.NoError(t, err, "Insert should not return an error")
	assert.Equal(t, 3, list.Size(), "Size should be 3 after insertion at tail")
	assert.Equal(t, 2, list.tail.val, "Tail value should be 3")
	assert.Equal(t, 3, list.tail.prev.val, "Previous value before tail should be 2")
	assert.Equal(t, 1, list.head.val, "Head value should remain 1")
}

func TestCDLL_Insert_Middle(t *testing.T) {
	list := createCDLLWithElements(1, 3)
	err := list.Insert(2, 1)
	require.NoError(t, err, "Insert should not return an error")
	assert.Equal(t, 3, list.Size(), "Size should be 3 after insertion in the middle")
	assert.Equal(t, 1, list.head.val, "Head value should remain 1")
	assert.Equal(t, 2, list.head.next.val, "Next value after head should be 2")
	assert.Equal(t, 3, list.tail.val, "Tail value should remain 3")
}

func TestCDLL_Get(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	val, err := list.Get(1)
	require.NoError(t, err, "Get should not return an error")
	assert.Equal(t, 2, *val, "Get should return the correct value")
}

func TestCDLL_Get_OutOfBounds(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	_, err := list.Get(3)
	assert.Error(t, err, "Get should return an error for out of bounds index")
}

func TestCDLL_Clear(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	err := list.Clear()
	require.NoError(t, err, "Clear should not return an error")
	assert.Equal(t, 0, list.Size(), "Size should be 0 after clear")
	assert.Nil(t, list.head, "Head should be nil after clear")
	assert.Nil(t, list.tail, "Tail should be nil after clear")
}

func TestCDLL_Contains(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	assert.True(t, list.Contains(2), "Contains should return true for an existing element")
	assert.False(t, list.Contains(4), "Contains should return false for a non-existing element")
}

func TestCDLL_Set(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	err := list.Set(4, 1)
	require.NoError(t, err, "Set should not return an error")
	val, _ := list.Get(1)
	assert.Equal(t, 4, *val, "Set should update the value at the given position")
}

func TestCDLL_GetPosition(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	pos, err := list.GetPosition(2)
	require.NoError(t, err, "GetPosition should not return an error")
	assert.Equal(t, 1, pos, "GetPosition should return the correct position of the element")
}

func TestCDLL_GetPosition_NotFound(t *testing.T) {
	list := createCDLLWithElements(1, 2, 3)
	_, err := list.GetPosition(4)
	assert.Error(t, err, "GetPosition should return an error if the element is not found")
}
