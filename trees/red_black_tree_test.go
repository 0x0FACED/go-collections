package trees

import (
	"fmt"
	"testing"
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

	tr.Insert(10)
	tr.Insert(9)
	tr.Insert(8)
	tr.Insert(7)
	tr.Insert(6)
	tr.Insert(3)
	tr.Insert(15)
	fmt.Println(tr.root.right.val)
	tr.Insert(11)

	fmt.Println(tr.InOrder())
	fmt.Println(tr.root.val)
	fmt.Println(tr.root.left.val)
	fmt.Println(tr.root.right.val)
	tr.PrintTree()
}
