package trees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST_Insert(t *testing.T) {
	compare := func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return -1
		} else {
			return 1
		}
	}

	tree := NewBST(compare)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(-1)

	val, err := tree.Search(5)
	assert.Error(t, err, "There is no value 5 in the tree")
	assert.Nil(t, val, "There is no 5, so must be val = nil")

	val, err = tree.Search(4)
	assert.NoError(t, err)
	assert.Equal(t, 4, *val)
}

func TestBST_Delete(t *testing.T) {
	compare := func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return -1
		} else {
			return 1
		}
	}

	tree := NewBST(compare)
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(70)
	tree.Insert(60)
	tree.Insert(80)

	err := tree.Delete(51)
	assert.Error(t, err, "There is no value 51 in the tree")

	err = tree.Delete(20)
	assert.NoError(t, err, "No error, just deleted the leaf")

	err = tree.Delete(50)
	assert.NoError(t, err, "No error, delete the root")

	root := tree.root.val
	assert.Equal(t, 60, root, "Must be 60, because this is the min of right")
}

func TestBST_Traversal(t *testing.T) {
	compare := func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return -1
		} else {
			return 1
		}
	}

	var tree TraversalTree[int]
	tree = NewBST(compare)
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(70)
	tree.Insert(60)
	tree.Insert(80)

	items1 := tree.InOrder()
	items2 := tree.PreOrder()
	items3 := tree.PostOrder()
	items4 := tree.LevelOrder()

	assert.Equal(t, []int{20, 30, 40, 50, 60, 70, 80}, items1)
	assert.Equal(t, []int{50, 20, 30, 40, 60, 70, 80}, items2)
	assert.Equal(t, []int{20, 30, 40, 60, 70, 80, 50}, items3)
	assert.Equal(t, []int{50, 30, 70, 20, 40, 60, 80}, items4)

	fmt.Println(items1)
	fmt.Println(items2)
	fmt.Println(items3)
	fmt.Println(items4)
}
