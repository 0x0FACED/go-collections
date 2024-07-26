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
	err := tr.Delete(123)
	assert.Error(t, err)

	// delete node with no children
	tr.Delete(93)

	fmt.Println(tr.InOrder())
	assert.Equal(t, COLOR(red), tr.root.left.clr)
	assert.Equal(t, COLOR(red), tr.root.right.clr)

	assert.Equal(t, COLOR(black), tr.root.left.right.clr)
	assert.Equal(t, COLOR(black), tr.root.right.right.clr)
	assert.Equal(t, COLOR(black), tr.root.left.left.clr)

	// delete node with 2 children
	val, err := tr.Search(95)
	assert.NoError(t, err)
	assert.Equal(t, 95, *val)

	// delete
	tr.Delete(95)

	fmt.Println(tr.InOrder())
	fmt.Println(tr.root.val)
	fmt.Println(tr.root.left.val)
	fmt.Println(tr.root.left.left.val)
	fmt.Println(tr.root.right.val)

}
