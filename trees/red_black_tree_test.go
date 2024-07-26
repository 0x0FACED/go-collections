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

	fmt.Println(tr.root.right.val)
	tr.Insert(11)

	fmt.Println(tr.InOrder())
	fmt.Println(tr.root.val)
	fmt.Println(tr.root.left.val)
	fmt.Println(tr.root.right.val)
	tr.PrintTree()
}
