package trees

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

// rbt - Red-Black Tree
type rbt[T comparable] struct {
	root *rbt_node[T]

	compare Comparator[T]
}

func NewRBT[T comparable](compare Comparator[T]) *rbt[T] {
	return &rbt[T]{compare: compare}
}

func (rbt *rbt[T]) Insert(item T) {
	newNode := rbt.insertHelper(rbt.root, item, nil)

	rbt.fixInsert(newNode)
}

func (rbt *rbt[T]) Delete(item T) error {
	node := rbt.searchHelper(rbt.root, item)
	if node == nil {
		return fmt.Errorf(gocollections.ErrNotFound)
	}

	if node.left == nil && node.right == nil {
		if node == node.parent.left {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		return nil
	}

	if (node.left != nil && node.right == nil) ||
		(node.left == nil && node.right != nil) {

		// TODO: delete when only 1 child
		return nil
	}

	// TODO: delete when 2 children
	return nil
}

func (rbt *rbt[T]) Search(item T) (*T, error) {
	node := rbt.searchHelper(rbt.root, item)
	if node != nil {
		return &node.val, nil
	}
	return nil, fmt.Errorf(gocollections.ErrNotFound)

}

func (rbt *rbt[T]) InOrder() []T {
	var items []T
	rbt.inOrderHelper(rbt.root, &items)
	return items
}

func (rbt *rbt[T]) PrintTree() {
	printTree(rbt.root, "", true)
}
