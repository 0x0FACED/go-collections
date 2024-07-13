package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList_Add(t *testing.T) {
	list := NewSinglyLinked[int]()
	list.Add(13)
	list.Add(15)
	list.Add(18)
	list.Add(20)
	list.Print()
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	list := NewSinglyLinked[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Print()
	actual, _ := list.Get(2)
	assert.Equal(t, 3, *actual)

	list.Insert(15, 2)
	actual2, _ := list.Get(3)
	assert.Equal(t, 3, *actual2)
	list.Print()

	list.Insert(50, 0)
	list.Print()
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	list := NewSinglyLinked[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Insert(15, 2)
	list.Insert(50, 0)
	list.Print()

	err := list.RemoveLast()
	assert.NoError(t, err)
	list.Print()

	actual, _ := list.Get(4)
	assert.Equal(t, 3, *actual)

	list.RemoveLast()
	list.RemoveLast()
	list.RemoveLast()
	list.RemoveLast()
	list.RemoveLast()
	list.RemoveLast()
	list.RemoveLast()
	list.RemoveLast()
	err = list.RemoveLast()
	assert.Error(t, err)
	list.Print()

	list.Add(123)
	list.Add(1)
	list.Add(12)
	list.Add(1232)
	list.Print()
	list.RemoveVal(123)
	list.Print()
	pos, err := list.RemoveVal(54134)
	assert.Equal(t, -1, pos)
	assert.Error(t, err)

	pos, err = list.RemoveVal(1232)
	assert.Equal(t, 2, pos)
	assert.NoError(t, err)
	list.Print()
}
