package trees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var compare = func(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func TestRBT_Insert(t *testing.T) {
	tr := NewRBT[int](compare)

	tr.Insert(100)
	tr.Insert(99)
	tr.Insert(98)
	tr.Insert(97)
	tr.Insert(96)
	tr.Insert(95)
	tr.Insert(94)
	fmt.Println(tr.root.val)
	fmt.Println(tr.root.right.val)
	fmt.Println(tr.root.left.val)
	tr.Insert(93)
	fmt.Println()
	fmt.Println()
	fmt.Println(tr.InOrder())
	fmt.Println(tr.root.val)
	fmt.Println(tr.root.left.val)
	fmt.Println(tr.root.right.val)
}

func TestRBT_Delete(t *testing.T) {
	tr := NewRBT[int](compare)

	tr.Insert(100)
	tr.Insert(99)
	tr.Insert(98)
	tr.Insert(97)
	tr.Insert(96)
	tr.Insert(95)
	tr.Insert(94)
	tr.Insert(93)
	fmt.Println()
	err := tr.Delete(123)
	assert.Error(t, err)

	// delete node with no children
	tr.Delete(93)

	assert.Equal(t, COLOR(red), tr.root.left.clr)
	assert.Equal(t, COLOR(red), tr.root.right.clr)

	assert.Equal(t, COLOR(black), tr.root.left.right.clr)
	assert.Equal(t, COLOR(black), tr.root.right.right.clr)
	assert.Equal(t, COLOR(black), tr.root.left.left.clr)

	// delete node with 2 children
	val, err := tr.Search(95)
	assert.NoError(t, err)
	assert.Equal(t, 95, *val)
	tr.PrintTree()

	// delete
	tr.Delete(95)

	tr.PrintTree()

	tr.Insert(121)
	tr.Insert(130)
	tr.PrintTree()
	tr.Delete(99)
	tr.PrintTree()

	for i := 90; i != 78; i-- {
		tr.Insert(i)
	}
	tr.PrintTree()
	tr.Delete(97)
	tr.PrintTree()

	tr.Delete(85)
	tr.PrintTree()

	tr.Delete(96)
	tr.PrintTree()

	tr.Delete(94)
	tr.PrintTree()

	tr.Delete(84)
	tr.PrintTree()

	tr.Delete(121)
	tr.PrintTree()

	tr.Delete(100)
	tr.PrintTree()

	assert.Equal(t, 89, tr.root.val)
	tr.Delete(98)
	tr.PrintTree()
	assert.Equal(t, 83, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 82, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 87, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 86, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 81, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 80, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 89, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 88, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 90, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 79, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Equal(t, 130, tr.root.val)

	tr.Delete(tr.root.val)
	tr.PrintTree()
	assert.Nil(t, tr.root)

}

func TestRBT_Delete_Extended(t *testing.T) {
	tr := NewRBT[int](compare)

	tr.Insert(50)
	tr.Insert(25)
	tr.Insert(75)
	tr.Insert(10)
	tr.Insert(30)
	tr.Insert(60)
	tr.Insert(80)
	tr.Insert(5)
	tr.Insert(15)
	tr.Insert(27)
	tr.Insert(55)
	tr.Insert(65)
	tr.Insert(70)
	tr.Insert(90)

	fmt.Println("Initial tree:")
	tr.PrintTree()

	err := tr.Delete(100)
	assert.Error(t, err, "Expected error for deleting non-existent node")

	tr.Delete(5)
	fmt.Println("After deleting leaf (5):")
	tr.PrintTree()

	tr.Delete(15)
	fmt.Println("After deleting node with one child (15):")
	tr.PrintTree()

	tr.Delete(25)
	fmt.Println("After deleting node with two children (25):")
	tr.PrintTree()

	tr.Delete(50)
	fmt.Println("After deleting the root (50):")
	tr.PrintTree()

	tr.Delete(30)
	fmt.Println("After deleting node with two children (30):")
	tr.PrintTree()

	tr.Insert(100)
	tr.Insert(110)
	tr.Insert(95)
	tr.PrintTree()

	tr.Delete(75)
	fmt.Println("After deleting node with two children (75):")
	tr.PrintTree()

	tr.Delete(90)
	fmt.Println("After deleting node with one child (90):")
	tr.PrintTree()

	tr.Delete(65)
	fmt.Println("After deleting node with two children (65):")
	tr.PrintTree()

	tr.Delete(27)
	fmt.Println("After deleting leaf (27):")
	tr.PrintTree()

	inOrder := tr.InOrder()
	expectedOrder := []int{10, 55, 60, 70, 80, 95, 100, 110}
	assert.Equal(t, expectedOrder, inOrder, "InOrder traversal does not match expected order")

	fmt.Println("Final tree:")
	tr.PrintTree()
	fmt.Println("InOrder traversal:", tr.InOrder())
}

func TestRBT_Traversal(t *testing.T) {
	tr := NewRBT[int](compare)

	tr.Insert(50)
	tr.Insert(25)
	tr.Insert(75)
	tr.Insert(10)
	tr.Insert(30)
	tr.Insert(60)
	tr.Insert(80)
	tr.Insert(5)
	tr.Insert(15)
	tr.Insert(27)
	tr.Insert(55)
	tr.Insert(65)
	tr.Insert(70)
	tr.Insert(90)

	fmt.Println("Initial tree:")
	tr.PrintTree()

	inOrder := tr.InOrder()
	preOrder := tr.PreOrder()
	postOrder := tr.PostOrder()
	levelOrder := tr.LevelOrder()
	fmt.Println("IN: ", inOrder)
	fmt.Println("PRE: ", preOrder)
	fmt.Println("POST: ", postOrder)
	fmt.Println("LEVEL: ", levelOrder)
	expectedInOrder := []int{5, 10, 15, 25, 27, 30, 50, 55, 60, 65, 70, 75, 80, 90}
	expectedPreOrder := []int{50, 25, 10, 5, 15, 30, 27, 75, 60, 55, 65, 70, 80, 90}
	expectedPostOrder := []int{5, 15, 10, 27, 30, 25, 55, 70, 65, 60, 90, 80, 75, 50}
	expectedLevelOrder := []int{50, 25, 75, 10, 30, 60, 80, 5, 15, 27, 55, 65, 90, 70}
	assert.Equal(t, expectedInOrder, inOrder, "inOrder traversal does not match expected order")
	assert.Equal(t, expectedPreOrder, preOrder, "preOrder traversal does not match expected order")
	assert.Equal(t, expectedPostOrder, postOrder, "postOrder traversal does not match expected order")
	assert.Equal(t, expectedLevelOrder, levelOrder, "levelOrder traversal does not match expected order")
}
