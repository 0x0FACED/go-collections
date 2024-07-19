package trees

import (
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
